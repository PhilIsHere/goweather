package jsonhandling

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var client *http.Client

func GenerateURL(apiLink string, apiEndpoint string, params map[string]string) (string, error) {
	baseURL, err := url.Parse(apiLink)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %v", err)
	}

	baseURL.Path += apiEndpoint

	parameters := url.Values{}
	for key, value := range params {
		parameters.Add(key, value)
	}
	baseURL.RawQuery = parameters.Encode()

	return baseURL.String(), nil
}

func CreateParams(params ...string) (map[string]string, error) {
	if len(params)%2 != 0 {
		return nil, fmt.Errorf("required even number of parameters, example: CreateParams(\"key1\", \"value1\", \"key2\", \"value2\")")
	}

	paramMap := make(map[string]string)
	for i := 0; i < len(params); i += 2 {
		paramMap[params[i]] = params[i+1]
	}

	return paramMap, nil
}

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
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)

	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return nil
}
