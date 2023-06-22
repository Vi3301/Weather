package main

import (
	"fmt"
	"io"
	"net/http"
)

type Сity struct {
	lat  float64
	long float64
}

func WeatherApiRequest(c Сity) (string, error) {
	url := fmt.Sprintf(`https://api.openweathermap.org/data/3.0/onecall?lat=%v&lon=%v&appid={"API_Key"}`, c.lat, c.long)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	// узнать что за Body
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	json.Unmarshal(body, &data)

}
func main() {

}
