package http

import (
	"encoding/json"
	"log"
	"net/http"
)

// DecodeResponse takes an HTTP Response in JSON format and decodes the
// Body into an interface type.
// Return the interface type.
func DecodeResponse(response *http.Response, v interface{}) *interface{} {
	if error := json.NewDecoder(response.Body).Decode(&v); error != nil {
		log.Println(error)
	}

	return &v
}
