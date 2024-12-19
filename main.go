package main

import (
	"net/http"

	"github.com/BartekDe/devops-thang/handlers"
)

func main() {
	http.HandleFunc("/version", handlers.VersionHandler)
	http.HandleFunc("/temperature", handlers.TemperatureHandler)

	http.ListenAndServe(":8080", nil)
}
