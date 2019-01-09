package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/seintun/dinesty.ninja-backend/config"
	. "github.com/seintun/dinesty.ninja-backend/dao"
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = BizDAO{}

// GetBizYelp POST & YELP GET JSON API
func GetBizYelp(w http.ResponseWriter, r *http.Request) {
	rB, _ := ioutil.ReadAll(r.Body)
	var yID YelpID
	err := json.Unmarshal([]byte(rB), &yID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	config.Read()
	yelpURL := config.YelpURL + yID.BusinessID
	bearer := "Bearer " + config.YelpKey
	biz, err := dao.FetchYelpJSN(yelpURL, bearer)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid YelpBiz ID")
		return
	}
	respondWithJson(w, http.StatusOK, biz)
}

// RegisterBiz insert new biz
func RegisterBiz(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var biz Biz
	if err := json.NewDecoder(r.Body).Decode(&biz); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	biz.ID = bson.NewObjectId()
	if err := dao.RegisterBiz(biz); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, biz)
}

// FetchBiz return list of all biz
func FetchBiz(w http.ResponseWriter, r *http.Request) {
	bizs, err := dao.FetchBiz()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, bizs)
}

// FindBiz return specified biz
func FindBizByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	biz, err := dao.FindBizByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Biz ID")
		return
	}
	respondWithJson(w, http.StatusOK, biz)
}

// UpdateBizByID update specified biz
func UpdateBizByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var biz Biz
	if err := json.NewDecoder(r.Body).Decode(&biz); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	params := mux.Vars(r)
	err := dao.UpdateBizByID(params["id"], biz)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Biz ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// // DeactivateBizByID
// func DeactivateBizByID(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	err := dao.DeactivateBizByID(params["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid Biz ID")
// 		return
// 	}
// 	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
// }

// DeleteBiz delete biz
func DeleteBizByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dao.DeleteBizByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Biz ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// respondWithError will identify error msg and respond back to the client
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// respondWithJson will Marshal and send response back to the client
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}
