package handlers

import (
	"fmt"
	"net/http"
)

func CreateObd2Data(http.ResponseWriter, *http.Request) {
	fmt.Println("Create Data")
}

func ShowObd2Data(http.ResponseWriter, *http.Request) {
	fmt.Println("Show Data")
}
