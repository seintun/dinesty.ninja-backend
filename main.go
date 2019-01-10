package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	ctrl "github.com/seintun/dinesty.ninja-backend/controllers"
)

func main() {
	r := mux.NewRouter()

	// HandleFun to route RESTful endpoints for /biz (business)
	r.HandleFunc("/biz/validate", ctrl.GetBizYelp).Methods("POST")
	r.HandleFunc("/biz/register", ctrl.RegisterBiz).Methods("POST")
	r.HandleFunc("/biz", ctrl.FetchBiz).Methods("GET")
	r.HandleFunc("/biz/{id}", ctrl.FindBizByID).Methods("GET")
	r.HandleFunc("/biz/{id}", ctrl.UpdateBizByID).Methods("PUT")
	// r.HandleFunc("/biz/{id}", ctrl.DeactivateBizByID).Methods("PUT")
	r.HandleFunc("/biz/{id}", ctrl.DeleteBizByID).Methods("DELETE")

	// HandleFun to route RESTful endpoints for /biz/{id}/menu (business's menu)
	r.HandleFunc("/biz/{id}/importmenu", ctrl.ImportJSON).Methods("POST")
	r.HandleFunc("/biz/{id}/menu", ctrl.CreateItem).Methods("POST")
	r.HandleFunc("/biz/{id}/menu", ctrl.FetchItems).Methods("GET")
	r.HandleFunc("/biz/{id}/menu/{mid}", ctrl.FindItemByID).Methods("GET")
	r.HandleFunc("/biz/{id}/menu/{mid}", ctrl.UpdateItemByID).Methods("PUT")
	r.HandleFunc("/biz/{id}/menu/{mid}", ctrl.DeleteItemByID).Methods("DELETE")

	// HandleFun to route RESTful endpoints for /users (users)
	r.HandleFunc("/users", ctrl.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", ctrl.FindUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", ctrl.UpdateUserByID).Methods("PUT")
	r.HandleFunc("/users/{id}", ctrl.DeleteUserByID).Methods("DELETE")

	// HandleFun to route RESTful endpoints for /orders (business and users orders)
	r.HandleFunc("/orders", ctrl.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", ctrl.FindOrderByID).Methods("GET")
	r.HandleFunc("/orders/users/{id}", ctrl.FetchOrdersByuserID).Methods("GET")
	r.HandleFunc("/orders/biz/{id}", ctrl.FetchOrdersBybizID).Methods("GET")
	r.HandleFunc("/orders/{id}", ctrl.UpdateOrderByID).Methods("PUT")
	r.HandleFunc("/orders/{id}/cancel", ctrl.CancelOrderByID).Methods("PUT")
	r.HandleFunc("/orders/{id}", ctrl.DeleteOrderByID).Methods("DELETE")
	r.HandleFunc("/orders/{id}/cart", ctrl.AddItemtoCart).Methods("PUT")
	r.HandleFunc("/orders/{id}/cart/{cid}", ctrl.DeleteItemfromCart).Methods("PUT")

	port := ":8080"
	log.Println("# Welcome to Dinesty.ninja, where our ninjas can infiltrate the wait to dine stealthly and rapidly! #")
	log.Println("###   A ninja(忍者 or 忍び) is listening your honorable commands on", port)
	rLog := handlers.LoggingHandler(os.Stdout, r)
	if err := http.ListenAndServe(":8080", rLog); err != nil {
		log.Fatal(err)
	}
}
