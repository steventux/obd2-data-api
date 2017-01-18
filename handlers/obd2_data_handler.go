package handlers

import (
	"github.com/steventux/obd2-data-api/helpers"
	"github.com/steventux/obd2-data-api/services"
	"net/http"
)

func CreateObd2Data(w http.ResponseWriter, r *http.Request) {
	if !helpers.Authenticate(r.FormValue("id")) {
		renderError(w, 403, "Forbidden")
		return
	}

	obd2Data := helpers.BuildObd2Data(r)
	_, err := services.SaveObd2Data(obd2Data)

	if err != nil {
		renderError(w, 500, "Couldn't save obd2 data")
		return
	}

	w.Write([]byte("OK"))
}

func ShowObd2Data(http.ResponseWriter, *http.Request) {
}

func renderError(w http.ResponseWriter, status int, errorString string) {
	w.WriteHeader(status)
	w.Write([]byte(errorString))
}
