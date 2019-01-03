package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/biz/validate", getBizYelp).Methods("POST")

	routerLogger := handlers.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":8080", routerLogger)
}
