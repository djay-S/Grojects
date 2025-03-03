package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	url    string = "https://api.openweathermap.org/data/2.5/weather?q=<city>&APPID=<api_id>"
	apiKey string = ""
)

type WeatherResp struct {
	Coordinates Coordinates   `json:"coord"`
	Weather     []WeatherInfo `json:"weather"`
	Base        string        `json:"base"`
	Main        WeatherAttr   `json:"main"`
	Visibility  int           `json:"visibility"`
	Wind        Wind          `json:"wind"`
	Clouds      Clouds        `json:"clouds"`
	Dt          int           `json:"dt"`
	Sys         Sys           `json:"sys"`
	Timezone    int           `json:"timezone"`
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Cod         int           `json:"cod"`
}

type WeatherInfo struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherAttr struct {
	Temp        float32 `json:"temp"`
	FeelsLike   float32 `json:"feels_like"`
	TempMin     float32 `json:"temp_min"`
	TempMax     float32 `json:"temp_max"`
	Pressure    float32 `json:"pressure"`
	Humidity    float32 `json:"humidity"`
	SeaLevel    float32 `json:"sea_level"`
	GroundLevel float32 `json:"grnd_level"`
}

type Wind struct {
	Speed float32 `json:"speed"`
	Deg   float32 `json:"deg"`
	Gust  float32 `json:"gust,omitempty"` // Optional field
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type Coordinates struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

//func main() {
//	log.Printf("Starting weather api server at: %v", time.Now())
//	log.Printf("Api key: %v", apiKey)
//	getWeatherData()
//}

func getWeatherData() {
	apiKey := getApiKey()
	url = strings.ReplaceAll(url, "<api_id>", apiKey)
	log.Printf("UrL: %v", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error occurred in fetching the weather details: %v", err)
		return
	}

	defer resp.Body.Close()

	log.Printf("Response: %v", resp)
	log.Printf("Response Body: %v", resp.Body)
	bodyResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occurred in reading the response: %v", err)
		return
	}

	var weatherResponse WeatherResp

	// json.NewDecoder(resp.Body).Decode(&weatherResponse)
	json.Unmarshal(bodyResponse, weatherResponse)
	log.Printf("Decoded response: %v", weatherResponse)
	log.Printf("Bosy string: %v", string(bodyResponse))
}

func GetWeatherDataResponse(cityName string) (string, error) {
	log.Printf("GetWeatherDataResponse for city: %v", cityName)
	apiKey := getApiKey()
	locUrl := strings.ReplaceAll(url, "<api_id>", apiKey)
	locUrl = strings.ReplaceAll(locUrl, "<city>", cityName)

	// log.Printf("Url %v", locUrl)

	resp, err := http.Get(locUrl)
	if err != nil {
		log.Printf("Error occurred %v", err)
		return "", err
	}
	defer resp.Body.Close()
	log.Printf("Response: %v", resp)

	bodyResponse, err := io.ReadAll(resp.Body)
	// log.Printf("Bosy string: %v", string(bodyResponse))
	if err != nil {
		log.Printf("Error occurred in reading the response: %v", err)
		return "", err
	}
	if resp.StatusCode != 200 {
		return string(bodyResponse), fmt.Errorf(string(bodyResponse))
	}
	return string(bodyResponse), nil
}

func getApiKey() string {
	if len(apiKey) > 0 {
		return apiKey
	}

	filePath := "api.txt"

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Printf("Error occurred in fetching api key: %v", err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines[0]
}
