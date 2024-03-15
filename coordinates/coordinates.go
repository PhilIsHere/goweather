package coordinates

import (
	"fmt"
	"goweather/jsonhandling"
	"net/url"
)

// Struct for the LAT and LON coordinates
type Location struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// Function to get the LAT and LON coordinates from the given city
func GetCoordinates(city string) (string, string, error) {
	//Initialize the URL
	baseURL, err := url.Parse("https://nominatim.openstreetmap.org/search")
	if err != nil {
		return "","",fmt.Errorf("error parsing URL: %v", err)
	}

	parameters := url.Values{}
	parameters.Add("q", city)
	parameters.Add("format", "jsonv2")
	parameters.Add("limit", "1")
	baseURL.RawQuery = parameters.Encode()


	//Initialize the location struct
	var location []Location

	//Get the JSON data from the API
	err = jsonhandling.GetJson(baseURL.String(), &location)
	if err != nil {
		return "","",fmt.Errorf("error getting coordinates: %v", err)
	}

	if len(location) == 0 {
		return "","",fmt.Errorf("no coordinates found for %s", city)
	}

	//Return the LAT and LON coordinates
	return location[0].Lat, location[0].Lon, nil
}
