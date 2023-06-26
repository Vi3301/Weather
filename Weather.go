package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type City struct {
	Lat  float64
	Long float64
}
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

func WeatherApiRequest(c City) (WeatherResponse, error) {
	apikey := os.Getenv("WEATHER_APITOKEN")
	url := fmt.Sprintf(`https://api.weather.yandex.ru/v2/forecast?lat=%v&lon=%v`, c.Lat, c.Long)
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
	var volgograd City = City{Lat: 48.71939, Long: 44.50183}
	WeatherApiRequest(volgograd)
}
