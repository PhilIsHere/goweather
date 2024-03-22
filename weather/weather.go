package weather

import (
	"fmt"
	"goweather/jsonhandling"
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
func GetWeather(jh jsonhandling.DefaultJSONHandler, lat string, lon string) (WeatherData, error) {

	//Parameters for GET request
	params, err := jh.CreateParams("lat", lat, "lon", lon)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error creating parameters: %v", err)
	}

	//Initialize the URL
	url, err := jh.GenerateURL("https://api.brightsky.dev", "current_weather", params)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error generating Weather-URL: %v", err)
	}

	//Initialize the weather struct
	var weather WeatherData

	//Get the JSON data from the API
	err = jh.GetJson(url, &weather)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error getting weather data: %v", err)
	}

	//Check if array is empty
	if len(weather.Sources) == 0 {
		return WeatherData{}, fmt.Errorf("no weather data found for %s, %s", lat, lon)
	}

	//Return the weather data
	return weather, nil
}

// Function to get the alerts from the given City
func GetAlerts(jh jsonhandling.DefaultJSONHandler, lat string, lon string) (Alerts, error) {

	//Parameters for GET request
	params, err := jh.CreateParams("lat", lat, "lon", lon)
	if err != nil {
		return Alerts{}, fmt.Errorf("error creating parameters: %v", err)
	}

	//Initialize the URL
	url, err := jh.GenerateURL("https://api.brightsky.dev", "alerts", params)
	if err != nil {
		return Alerts{}, fmt.Errorf("error generating Alerts-URL: %v", err)
	}

	//Initialize the alerts struct
	var alerts Alerts

	//Get the JSON data from the API
	jh.GetJson(url, &alerts)

	//Return the alerts
	return alerts, nil
}
