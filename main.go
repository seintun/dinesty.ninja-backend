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
	r.HandleFunc("/biz/register", ctrl.RegisterBiz).Methods("POST")
	r.HandleFunc("/biz", ctrl.FetchBiz).Methods("GET")
	r.HandleFunc("/biz/{id}", ctrl.FindBizByID).Methods("GET")
	r.HandleFunc("/biz/{id}", ctrl.UpdateBizByID).Methods("PUT")
	// r.HandleFunc("/biz/{id}", ctrl.DeactivateBizByID).Methods("PUT")
	r.HandleFunc("/biz/{id}", ctrl.DeleteBizByID).Methods("DELETE")

	r.HandleFunc("/biz/{id}/menu", ctrl.CreateItem).Methods("POST")
	r.HandleFunc("/biz/{id}/menu", ctrl.FetchItems).Methods("GET")

	rLog := handlers.LoggingHandler(os.Stdout, r)
	if err := http.ListenAndServe(":8080", rLog); err != nil {
		log.Fatal(err)
	}
}
