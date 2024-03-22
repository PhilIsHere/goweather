package coordinates

import (
	"fmt"
	"goweather/jsonhandling"
)

// Struct for the LAT and LON coordinates
type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// Function to get the LAT and LON coordinates from the given city
func GetCoordinates(jh jsonhandling.DefaultJSONHandler, userInput string) (string, string, error) {

	//Parameters for GET request
	params, err := jh.CreateParams("q", userInput, "format", "jsonv2", "limit", "1")
	if err != nil {
		return "", "", fmt.Errorf("error creating parameters: %v", err)
	}

	//Initialize the URL
	url, err := jh.GenerateURL("https://nominatim.openstreetmap.org", "search", params)
	if err != nil {
		return "", "", fmt.Errorf("error generating URL: %v", err)
	}

	//Initialize the location struct
	var location []Location

	//Get the JSON data from the API
	err = jh.GetJson(url, &location)
	if err != nil {
		return "", "", fmt.Errorf("error getting coordinates: %v", err)
	}

	if len(location) == 0 {
		return "", "", fmt.Errorf("no coordinates found for %s", userInput)
	}

	//Return the LAT and LON coordinates
	return location[0].Lat, location[0].Lon, nil
}
