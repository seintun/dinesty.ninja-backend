package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/seintun/dinesty.ninja-backend/config"
	. "github.com/seintun/dinesty.ninja-backend/dao"
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = BizDAO{}

// RegisterBiz POST
func RegisterBiz(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var biz Biz
	if err := json.NewDecoder(r.Body).Decode(&biz); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	biz.ID = bson.NewObjectId()
	if err := dao.Insert(biz); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, biz)
}

// GetBizYelp POST & YELP GET
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

	request, _ := http.NewRequest("GET", yelpURL, nil)
	request.Header.Add("Authorization", bearer)
	client := &http.Client{}
	yelpR, _ := client.Do(request)
	data, _ := ioutil.ReadAll(yelpR.Body)
	defer yelpR.Body.Close()

	var yJSN YelpJSN
	json.Unmarshal([]byte(data), &yJSN)
	respondWithJson(w, http.StatusCreated, yJSN)
}

// GET list of movies
func FetchBiz(w http.ResponseWriter, r *http.Request) {
	bizs, err := dao.FetchBiz()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, bizs)
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
