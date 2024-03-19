/*
This is a simple program that uses the https://brightsky.dev API to get the weather for a given uinput.
To recieve the needed LAT and LON coordinates, the program uses the https://nominatim.openstreetmap.org API.
The temperature is displayed in both Celsius and Fahrenheit.
*/

package main

import (
	"bufio"
	"fmt"
	"goweather/coordinates"
	"goweather/weather"
	"log"
	"os"
	"time"
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
	reader := bufio.NewReader(os.Stdin)
	inputChannel := make(chan string)

	//Ask the user for the city
	fmt.Println("Bitte gib eine deutsche Adresse, PLZ oder Sehenswürdigkeit ein: ")
	var uinput string
	uinput, _ = reader.ReadString('\n')

	//Get the LAT and LON coordinates
	lat, lon, err := coordinates.GetCoordinates(uinput)
	if err != nil {
		log.Fatal(err)
	}

	//Get Weather Data
	wData, err := weather.GetWeather(lat, lon)
	if err != nil {
		log.Fatal(err)
	}

	//Get Alerts
	wAlerts, err := weather.GetAlerts(lat, lon)
	if err != nil {
		log.Fatal(err)
	}

	//Print the weather data
	fmt.Println("Die Station", wData.Sources[0].StationName, "meldet für", uinput, wData.Weather.Condition, "bei", wData.Weather.Temperature, "°C")

	//Print the alerts red and bold if there are any
	if wAlerts.Alerts.EventDe != "" {
		fmt.Println("\033[1;31m ACHTUNG! Für", wAlerts.Location.Name, "gibt es eine Amtliche Unwetterwarnung! \033[0m")
		fmt.Println("Art:", wAlerts.Alerts.EventDe)
		fmt.Println("Beschreibung:", wAlerts.Alerts.DescriptionDe)
		fmt.Println("Betroffenes Gebiet:", wAlerts.Location.Name, wAlerts.Location.District, wAlerts.Location.State)
		fmt.Println("Gültig von", wAlerts.Alerts.Effective, "bis", wAlerts.Alerts.Expires)
		fmt.Println()
	} else {
		fmt.Println("Es gibt keine amtlichen Unwetterwarnungen")
		fmt.Println()
	}

	//Wait for user input to close the program or close after 5 seconds
	fmt.Println("Drücke eine beliebige Taste, um das Programm zu beenden.")
	go func() {
		go func() {
			for i := 5; i > 0; i-- {
				fmt.Printf("\rDas Programm wird in %d Sekunden automatisch beendet...", i-1)
				time.Sleep(1 * time.Second)
			}
		}()
		input, _ := reader.ReadString('\n')
		inputChannel <- input
	}()
	select {
	case <-inputChannel:
		log.Println("Programm beendet")
	case <-time.After(5 * time.Second):
	}
	os.Exit(0)
}
