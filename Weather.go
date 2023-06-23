package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type City struct {
	lat  float64
	long float64
}

func WeatherApiRequest(c City) (string, error) {
	apikey := os.Getenv("WEATHER_APITOKEN")
	url := fmt.Sprintf(`https://api.openweathermap.org/data/3.0/onecall?lat=%v&lon=%v&appid=%v`, c.lat, c.long, apikey)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return " ", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return " ", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return " ", err
	}
	defer resp.Body.Close()
	var data string
	json.Unmarshal(body, &data)
	fmt.Println(data)
	return data, nil
}
func main() {
	var olhovka City = City{lat: 33.44, long: -94.04}
	res, err := WeatherApiRequest(olhovka)
	fmt.Println(res)
	fmt.Println(err)
}
