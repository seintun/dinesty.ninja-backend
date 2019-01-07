package controllers

import (
	"encoding/json"
	"net/http"

	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

// CreateUser insert new business
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	u.ID = bson.NewObjectId()
	if err := dao.CreateUser(u); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, u)
}
