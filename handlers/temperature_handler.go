package handlers

import (
	"net/http"
)

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Temperature"))
}
