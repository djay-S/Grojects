package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type city struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	log.Printf("Starting gateway-server at")

	http.HandleFunc("/cities", getCities)

	http.HandleFunc("/weather/city/{cityName}", fetchWeatherData)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	//(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/cities")
}

func getCities(w http.ResponseWriter, r *http.Request) {
	handleCors(&w)
	log.Printf("Fetching cities")
	cities := []city{{Name: "Delhi", Country: "in"}, {Name: "Tokyo", Country: "jp"}}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cities)
}

func fetchWeatherData(w http.ResponseWriter, r *http.Request) {
	handleCors(&w)
	cityName := r.PathValue("cityName")
	log.Printf("Fetching weather details for city: %v", cityName)
	resp, err := GetWeatherDataResponse(cityName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(resp)
	fmt.Fprint(w, resp)
}
