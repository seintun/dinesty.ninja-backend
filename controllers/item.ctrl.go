package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

// CreateItem insert new business
func CreateItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	item.ID = bson.NewObjectId()
	if err := dao.CreateItem(item); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, item)
}

// ImportJSON passes array of menuItem objects JSON into dao.ImportJSON
func ImportJSON(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var iJSN []Item
	if err := json.NewDecoder(r.Body).Decode(&iJSN); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.ImportJSON(iJSN); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, map[string]string{"result": "success"})
}

// FetchItems return all menuItems by itemID
func FetchItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	items, err := dao.FetchItems(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid items ID")
		return
	}
	respondWithJson(w, http.StatusOK, items)
}

// FindItemByID return specified menuItem by itemID
func FindItemByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	biz, err := dao.FindItemByID(params["mid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Menu ID")
		return
	}
	respondWithJson(w, http.StatusOK, biz)
}

// UpdateItemByID update specified menuItem by itemID
func UpdateItemByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	params := mux.Vars(r)
	err := dao.UpdateItemByID(params["mid"], item)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteItemByID delete specified menuItem by itemID
func DeleteItemByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dao.DeleteItemByID(params["mid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
