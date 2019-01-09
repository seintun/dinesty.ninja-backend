package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

// CreateOrder insert new order
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var o Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	o.ID = bson.NewObjectId()
	if err := dao.CreateOrder(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, o)
}

// FindOrderByID return specified order
func FindOrderByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	o, err := dao.FindOrderByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	respondWithJson(w, http.StatusOK, o)
}

// DeleteOrderByID delete specified order
func DeleteOrderByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dao.DeleteOrderByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// AddItemtoMenu update specified menuItem by itemID
func AddItemtoMenu(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	params := mux.Vars(r)
	err := dao.AddItemtoMenu(params["id"], item)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
