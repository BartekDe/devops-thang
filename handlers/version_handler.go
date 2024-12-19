package handlers

import (
	"net/http"
)

const version = "0.0.1"

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(version))
}
