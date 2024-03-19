package weather

import (
	"fmt"
	"goweather/jsonhandling"
	"net/url"
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

//Function to generate the base URL for the Brightsky API

func generateURL(lat string, lon string, endpoint string) (string, error) {
	baseURL, err := url.Parse("https://api.brightsky.dev/")
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %v", err)
	}

	baseURL.Path += endpoint

	parameters := url.Values{}
	parameters.Add("lat", lat)
	parameters.Add("lon", lon)
	baseURL.RawQuery = parameters.Encode()

	return baseURL.String(), nil
}

// Function to get the weather data from the given City
func GetWeather(lat string, lon string) (WeatherData, error) {

	//Initialize the URL
	url, err := generateURL(lat, lon, "current_weather")
	if err != nil {
		return WeatherData{}, fmt.Errorf("error generating Weather-URL: %v", err)
	}

	//Initialize the weather struct
	var weather WeatherData

	//Get the JSON data from the API
	jsonhandling.GetJson(url, &weather)

	//Check if array is empty
	if len(weather.Sources) == 0 {
		return WeatherData{}, fmt.Errorf("no weather data found for %s, %s", lat, lon)
	}

	//Return the weather data
	return weather, nil
}

// Function to get the alerts from the given City
func GetAlerts(lat string, lon string) (Alerts, error) {

	//Initialize the URL
	url, err := generateURL(lat, lon, "alerts")
	if err != nil {
		return Alerts{}, fmt.Errorf("error generating Alerts-URL: %v", err)
	}

	//Initialize the alerts struct
	var alerts Alerts

	//Get the JSON data from the API
	jsonhandling.GetJson(url, &alerts)

	//Return the alerts
	return alerts, nil
}
