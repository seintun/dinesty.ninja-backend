package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

// CreateUser insert new user
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

// FindUserByID return specified user
func FindUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	u, err := dao.FindUserByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	respondWithJson(w, http.StatusOK, u)
}

// UpdateUserByID update specified user
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	params := mux.Vars(r)
	err := dao.UpdateUserByID(params["id"], u)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid u ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteUserByID delete specified user
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dao.DeleteUserByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Biz ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
