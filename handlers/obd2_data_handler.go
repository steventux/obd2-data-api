package handlers

import (
	"github.com/gorilla/mux"
	"github.com/nbio/httpcontext"
	"github.com/steventux/obd2-data-api/helpers"
	"github.com/steventux/obd2-data-api/services"
	"github.com/zbindenren/negroni-mongo"
	"gopkg.in/mgo.v2"
	"net/http"
)

func CreateObd2Data(w http.ResponseWriter, r *http.Request) {
	if !helpers.Authenticate(r.FormValue("id")) {
		renderError(w, 403, "Forbidden")
		return
	}

	obd2Data := helpers.BuildObd2Data(r)
	db := getDatabaseFromRequestContext(r)
	_, err := services.SaveObd2Data(db, obd2Data)

	if err != nil {
		renderError(w, 500, "Couldn't save obd2 data")
		return
	}

	w.Write([]byte("OK"))
}

func ShowObd2Data(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var session = vars["session"]
	if session == "" {
		// Get latest session
		db := getDatabaseFromRequestContext(r)
		session = services.GetLatestObd2DataSession(db)
	}
	// Get data by session and respond with it.
}

func renderError(w http.ResponseWriter, status int, errorString string) {
	w.WriteHeader(status)
	w.Write([]byte(errorString))
}

func getDatabaseFromRequestContext(r *http.Request) *mgo.Database {
	return httpcontext.Get(r, negronimongo.ContextKey).(*mgo.Database)
}
