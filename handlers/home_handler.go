package handlers

import (
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

var renderer = render.New(render.Options{})

func Home(w http.ResponseWriter, r *http.Request) {
	renderer.JSON(w, http.StatusOK, map[string]string{"status": "OK"})
}
