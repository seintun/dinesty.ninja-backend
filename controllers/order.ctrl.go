package controllers

import (
	"encoding/json"
	"net/http"

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
