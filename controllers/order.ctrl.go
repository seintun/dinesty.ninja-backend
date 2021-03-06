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

// FetchOrdersByuserID return specified order
func FetchOrdersByuserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	o, err := dao.FetchOrdersByuserID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	respondWithJson(w, http.StatusOK, o)
}

// FetchOrdersBybizID return specified order
func FetchOrdersBybizID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	o, err := dao.FetchOrdersBybizID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid business ID")
		return
	}
	respondWithJson(w, http.StatusOK, o)
}

// UpdateOrderByID update specified order
func UpdateOrderByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var o Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	params := mux.Vars(r)
	err := dao.UpdateOrderByID(params["id"], o)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	respondWithJson(w, http.StatusOK, o)
}

// CancelOrderByID will change the active bool to false
func CancelOrderByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dao.CancelOrderByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "Your order has been cancelled"})
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

// AddItemtoCart update specified menuItem by itemID
func AddItemtoCart(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	params := mux.Vars(r)
	err := dao.AddItemtoCart(params["id"], item)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteItemfromCart update specified menuItem by itemID
func DeleteItemfromCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dao.DeleteItemfromCart(params["id"], params["cid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid params ID")
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
