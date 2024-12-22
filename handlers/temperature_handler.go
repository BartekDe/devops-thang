package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type WeatherData struct {
	Temperature float64 `json:"temp_c"`
}

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	loc := r.URL.Query().Get("loc")
	if len(loc) == 0 {
		w.Write([]byte("loc query param is required"))
		return
	}

	temp, err := getTemperature(loc)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(strconv.FormatFloat(temp, 'f', -1, 64)))
}

func getTemperature(loc string) (float64, error) {
	apiKey, err := getApiKey()
	if err != nil {
		return 0.0, err
	}

	url := "https://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + loc

	response, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}
	defer response.Body.Close()

	var result struct {
		Current WeatherData `json:"current"`
	}

	json.Unmarshal(body, &result)

	if result.Current.Temperature == 0 {
		fmt.Println("temp was zero")
		return 0.0, nil
	}

	return result.Current.Temperature, nil
}

func getApiKey() (string, error) {
	apiKey := os.Getenv("API_KEY")
	fmt.Println("api key: " + apiKey)
	if len(apiKey) == 0 {
		return "", errors.New("API_KEY is not set")
	}

	return apiKey, nil
}
