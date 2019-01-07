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

	r.HandleFunc("/biz/validate", ctrl.GetBizYelp).Methods("POST")
	// r.HandleFunc("/biz", fetchBiz).Methods("GET")
	// r.HandleFunc("/biz/register", ctrl.RegisterBiz).Methods("POST")

	rLog := handlers.LoggingHandler(os.Stdout, r)
	if err := http.ListenAndServe(":8080", rLog); err != nil {
		log.Fatal(err)
	}
}
