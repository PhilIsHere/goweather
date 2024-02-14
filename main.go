/*
This is a simple program that uses the https://brightsky.dev API to get the weather for a given city.
To recieve the needed LAT and LON coordinates, the program uses the https://nominatim.openstreetmap.org API.
It also accepts postal codes.
The temperature is displayed in both Celsius and Fahrenheit.
*/

package main

import (
	"fmt"
	"goweather/coordinates"
	"goweather/weather"
	"log"
	"strings"
)

// Struct for the weather data
type WeatherData struct {
	Weather struct {
		Condition   string  `json:"condition"`
		Temperature float64 `json:"temperature"`
	} `json:"weather"`
	Sources []struct {
		StationName string `json:"station_name"`
	} `json:"sources"`
}

// Main function
func main() {

	//Ask the user for the city
	fmt.Println("Bitte gib eine Stadt, PLZ oder Ort von Interesse ein: ")
	var city string
	fmt.Scanln(&city)

	//Check for German UmLaute
	if strings.ContainsAny(city, "äöüß") {
		log.Fatal("Ersetze bitte die Umlaute ä, ö, ü und ß durch ae, oe, ue und ss und versuche es erneut.")
	}

	//Get the LAT and LON coordinates
	lat, lon := coordinates.GetCoordinates(city)

	//Get Weather Data
	wData := weather.GetWeather(lat, lon)

	//Get Alerts
	wAlerts := weather.GetAlerts(lat, lon)

	//Print the weather data
	fmt.Println("Die Station", wData.Sources[0].StationName, "meldet für", city, wData.Weather.Condition, "bei", wData.Weather.Temperature, "°C")

	//Print the alerts red and bold if there are any
	if wAlerts.Alerts.EventDe != "" {
		fmt.Println("\033[1;31m ACHTUNG! Für", wAlerts.Location.Name, "gibt es eine Amtliche Unwetterwarnung! \033[0m")
		fmt.Println("Art:", wAlerts.Alerts.EventDe)
		fmt.Println("Beschreibung:", wAlerts.Alerts.DescriptionDe)
		fmt.Println("Betroffenes Gebiet:", wAlerts.Location.Name, wAlerts.Location.District, wAlerts.Location.State)
		fmt.Println("Gültig von", wAlerts.Alerts.Effective, "bis", wAlerts.Alerts.Expires)
	} else {
		fmt.Println("Es gibt keine amtlichen Unwetterwarnungen")
	}
}
