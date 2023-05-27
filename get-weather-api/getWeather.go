package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponse struct {
	CurrentWeather WeatherData `json:"current_weather"`
}

type WeatherData struct {
	Temperature float64 `json:"temperature"`
	WeatherCode int     `json:"weathercode"`
}

func convertWeatherCode(code int) string {
	switch code {
	case 0:
		return "sunny"
	case 1:
		return "mainly clear"
	case 2:
		return "partly cloudy"
	case 3:
		return "overcast"
	}
	return "unknown"
}

func main() {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a GET request for the weather API
	req, err := http.NewRequest("GET", "https://api.open-meteo.com/v1/forecast?latitude=48.85&longitude=2.35&current_weather=true", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the JSON response
	var weatherResp WeatherResponse
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert the weather code to a string
	weather := convertWeatherCode(weatherResp.CurrentWeather.WeatherCode)

	// Print the desired data
	fmt.Printf("The weather in Paris is %s and the temperature is %.2f°C\n", weather, weatherResp.CurrentWeather.Temperature)
}
