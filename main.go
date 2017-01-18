package main

import (
	"github.com/gorilla/mux"
	"github.com/steventux/obd2-data-api/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/obd2-data/create", handlers.CreateObd2Data).Methods("GET")
	r.HandleFunc("/obd2-data/show", handlers.ShowObd2Data).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
