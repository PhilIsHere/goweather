package jsonhandling

import (
	"encoding/json"
	"net/http"
	"time"
)

var client *http.Client

// Function to get a JSON Object from the given URL
// The URL is the API endpoint and the target is the struct where the JSON data will be stored
func GetJson(url string, target interface{}) error {

	//Initialize the HTTP client with a 10 second timeout if the API is slow or unresponsive
	client = &http.Client{Timeout: 10 * time.Second}

	resp, error := client.Get(url)
	if error != nil {
		return error
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
