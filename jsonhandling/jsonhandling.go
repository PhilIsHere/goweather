package jsonhandling

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var client *http.Client

// Function to get a JSON Object from the given URL
// The URL is the API endpoint and the target is the struct where the JSON data will be stored
func GetJson(url string, target interface{}) error {

	//Initialize the HTTP client with a 5 second timeout if the API is slow or unresponsive
	client = &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	//Set the User-Agent
	request.Header.Set("User-Agent", "PhilIsHere/goweather")

	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return fmt.Errorf("error reading response: %v", err)

	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return nil
}
