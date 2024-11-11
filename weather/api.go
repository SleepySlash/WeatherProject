package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func WeatherByPlace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")

	place := mux.Vars(r)["location"]
	theUrl := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=3eb7808380a842e39b5131134241111&q=%s", place)

	response, err := http.Get(theUrl)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
	}

	var result map[string]interface{}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}

func WeatherByPlaceDays(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")

	place := mux.Vars(r)["location"]
	days := mux.Vars(r)["days"]
	theUrl := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=3eb7808380a842e39b5131134241111&q=%s&days=%s", place, days)

	response, err := http.Get(theUrl)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
	}

	var result map[string]interface{}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}

func WeatherByPlaceDaysAqi(w http.ResponseWriter, r *http.Request) {
	place := mux.Vars(r)["location"]
	days := mux.Vars(r)["days"]
	aqi := mux.Vars(r)["aqi"]

	if aqi != "" {
		aqi = "yes"
	} else {
		aqi = "no"
	}

	theUrl := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=3eb7808380a842e39b5131134241111&q=%s&days=%s&aqi=%s", place, days, aqi)

	response, err := http.Get(theUrl)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
	}

	var result map[string]interface{}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}
