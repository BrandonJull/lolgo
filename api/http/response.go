package http

import (
	"log"
	"net/http"
)

// NewResponse generates and returns a HTTP reponse given
// an HTTP API request.
func NewResponse(request *http.Request) *http.Response {
	client := &http.Client{}

	response, error := client.Do(request)
	if error != nil {
		log.Fatal("NewResponse: ", error)
		return nil
	}

	return response
}
