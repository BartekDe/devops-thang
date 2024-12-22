package main

import (
	"net/http"

	"github.com/BartekDe/devops-thang/handlers"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load("./.env")

	http.HandleFunc("/version", handlers.VersionHandler)
	http.HandleFunc("/temperature", handlers.TemperatureHandler)

	http.ListenAndServe(":8080", nil)
}
