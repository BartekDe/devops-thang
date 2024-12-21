package handlers

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestTemperatureHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/temperature?loc=warsaw", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	TemperatureHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	body := rr.Body.String()
	_, err = strconv.ParseFloat(body, 64)
	if err != nil {
		t.Errorf("not a float value in the body: %v", body)
	}
}

func TestLocParamNotProvided(t *testing.T) {
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

	if rr.Body.String() != "loc query param is required" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "loc query param is required")
	}
}
