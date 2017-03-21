package main

import (
	"github.com/gorilla/mux"
	"github.com/steventux/obd2-data-api/handlers"
	"github.com/urfave/negroni"
	"github.com/zbindenren/negroni-mongo"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/obd2-data/create", handlers.CreateObd2Data).Methods("GET")
	r.HandleFunc("/obd2-data/show", handlers.ShowObd2Data).Methods("GET")
	n := negroni.New()
	m := negronimongo.NewMongoMiddleware(os.Getenv("MONGODB_HOSTS") + "/" + os.Getenv("MONGODB_DB"))
	n.Use(m)
	n.UseHandler(r)
	log.Fatal(http.ListenAndServe(":8000", n))
}
