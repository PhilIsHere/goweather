package coordinates

import (
	"fmt"
	"goweather/jsonhandling"
	"os"
)

// Struct for the LAT and LON coordinates
type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// Function to get the LAT and LON coordinates from the given city
func GetCoordinates(city string) (string, string) {
	//Initialize the URL
	url := "https://nominatim.openstreetmap.org/search?q=" + city + "&format=jsonv2&limit=1"

	//Initialize the location struct
	var location []Location

	//Get the JSON data from the API
	resp := jsonhandling.GetJson(url, &location)

	//Check for errors
	if resp != nil {
		fmt.Println("Error getting coordinates: ", resp)
		os.Exit(1)
	}

	//Return the LAT and LON coordinates
	return location[0].Lat, location[0].Lon
}
