package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type WeatherResponse struct {
	Info struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"info"`
	Fact struct {
		Temp       float64 `json:"temp"`
		Feels_like float64 `json:"feels_like"`
	} `json:"fact"`
}

func GeoCode(address string) (float64, float64, error) {
	apikey := os.Getenv("GEOCODER_APITOKEN")
	geocodeURL := fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?geocode=%s&apikey=%s&format=json", url.QueryEscape(address), apikey)
	response, err := http.Get(geocodeURL)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to Get request")
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, 0, fmt.Errorf("failed reading body")
	}
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, 0, err
	}
	pos := data["response"].(map[string]interface{})["GeoObjectCollection"].(map[string]interface{})["featureMember"].([]interface{})[0].(map[string]interface{})["GeoObject"].(map[string]interface{})["Point"].(map[string]interface{})["pos"].(string)
	coordinates := strings.Split(pos, " ")
	latitude, _ := strconv.ParseFloat(coordinates[1], 64)
	longitude, _ := strconv.ParseFloat(coordinates[0], 64)

	return latitude, longitude, nil
}
func WeatherApiRequest(latitude, longitude float64) (WeatherResponse, error) {
	apikey := os.Getenv("WEATHER_APITOKEN")
	url := fmt.Sprintf(`https://api.weather.yandex.ru/v2/forecast?lat=%v&lon=%v`, latitude, longitude)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Yandex-API-Key", apikey)
	if err != nil {
		return WeatherResponse{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherResponse{}, err
	}
	var data WeatherResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("failed to parsed response body: %v", err)
	}
	fmt.Printf("Temperature equal = %v°\n", data.Fact.Temp)
	return data, nil
}
func main() {
	var city string
	fmt.Println("Please enter the city where you want to find out the weather!")
	fmt.Scanf(`%s`, &city)
	lat, lon, err := GeoCode(city)
	if err != nil {
		return
	}
	WeatherApiRequest(lat, lon)
}
