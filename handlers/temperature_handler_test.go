package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemperatureHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/temperature", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	TemperatureHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	if rr.Body.String() != "Temperature" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "Temperature")
	}
}
