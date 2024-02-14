package weather

import (
	"fmt"
	"goweather/jsonhandling"
	"os"
)

// struct for the weather data from brightsky API
type WeatherData struct {
	Weather struct {
		Condition   string  `json:"condition"`
		Temperature float64 `json:"temperature"`
	} `json:"weather"`
	Sources []struct {
		StationName string `json:"station_name"`
	} `json:"sources"`
}

// Struct for the alerts
type Alerts struct {
	Alerts struct {
		Effective     string `json:"effective"`
		Expires       string `json:"expires"`
		EventDe       string `json:"event_de"`
		HeadlineDe    string `json:"headline_de"`
		DescriptionDe string `json:"description_de"`
	} `json:"alerts"`
	Location struct {
		Name     string `json:"name"`
		District string `json:"district"`
		State    string `json:"state"`
	}
}

// Function to get the weather data from the given City
func GetWeather(lat string, lon string) WeatherData {

	//Initialize the URL
	url := "https://api.brightsky.dev/current_weather?lat=" + lat + "&lon=" + lon

	//Initialize the weather struct
	var weather WeatherData

	//Get the JSON data from the API
	jsonhandling.GetJson(url, &weather)

	//Check if array is empty
	if len(weather.Sources) == 0 {
		fmt.Println("Please enter a german city or POI")
		os.Exit(2)
	}

	//Return the weather data
	return weather
}

// Function to get the alerts from the given City
func GetAlerts(lat string, lon string) Alerts {

	//Initialize the URL
	url := "https://api.brightsky.dev/alerts?lat=" + lat + "&lon=" + lon

	//Initialize the alerts struct
	var alerts Alerts

	//Get the JSON data from the API
	jsonhandling.GetJson(url, &alerts)

	//Return the alerts
	return alerts
}
