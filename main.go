package main

import (
	"net/http"
)

func main() {
	const version = "0.0.1"

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	})

	http.ListenAndServe(":8080", nil)
}
