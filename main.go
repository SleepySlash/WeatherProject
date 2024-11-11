package main

import (
	"net/http"
	"weather_REST/weather"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/weather/{location}", weather.WeatherByPlace).Methods("GET")
	router.HandleFunc("/weather/{location}/{days}", weather.WeatherByPlaceDays).Methods("GET")
	router.HandleFunc("/weather/{location}/{days}/{aqi}", weather.WeatherByPlaceDaysAqi).Methods("GET")
	http.ListenAndServe(":3000", router)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "Application/json")
	w.Write([]byte("Running on the port 3000"))
}
