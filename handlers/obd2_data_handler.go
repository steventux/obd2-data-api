package handlers

import (
	//"github.com/steventux/obd2-data-api/helpers"
	//"github.com/steventux/obd2-data-api/services"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

var rd = render.New()

func CreateObd2Data(w http.ResponseWriter, r *http.Request) {
	//obd2Data := helpers.BuildObd2Data(r)
	//_, err := services.SaveObd2Data(obd2Data)

	//if err != nil {
	//	renderError(w, 500, "Couldn't save obd2 data")
	//}

	rd.JSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

func ShowObd2Data(http.ResponseWriter, *http.Request) {
}

func renderError(w http.ResponseWriter, status int, errorString string) {
	rd.JSON(w, status, map[string]string{"error": errorString})
}
