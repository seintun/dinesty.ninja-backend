package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

// RegisterBiz insert new business
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

// FetchItems all menuItems by bizID
func FetchItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	items, err := dao.FetchItems(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid items ID")
		return
	}
	respondWithJson(w, http.StatusOK, items)
}

// FindItem menuItem by bizID
func FindItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	biz, err := dao.FindItem(params["mid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Menu ID")
		return
	}
	respondWithJson(w, http.StatusOK, biz)
}
