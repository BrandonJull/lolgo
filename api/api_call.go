package api

import (
	"log"
	"lolgo/api/http"
)

// NewAPICall creates a new call the Riot API.
// Accepts a URL for the specific call to the API and interface type.
// Return the interface type after JSON has been decoded.
func NewAPICall(url string, v interface{}) *interface{} {
	// Create a new HTTP request to the Riot API given our URL.
	request := http.NewRequest(url)

	// Capture the HTTP response.
	response := http.NewResponse(request)

	// Defer closing the response body.
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log.Println("API response returned error:", response.Status, "| Request:", request.URL)
		return nil
	}

	// Decode the response.
	http.DecodeResponse(response, &v)

	return &v
}
