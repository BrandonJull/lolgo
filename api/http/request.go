package http

import (
	"log"
	"net/http"
)

// NewRequest generates and returns a HTTP request given
// a generated URL call to the Riot API.
func NewRequest(url string) *http.Request {
	request, error := http.NewRequest("GET", url, nil)

	if error != nil {
		log.Fatal("NewRequest: ", error)
		return nil
	}

	return request
}
