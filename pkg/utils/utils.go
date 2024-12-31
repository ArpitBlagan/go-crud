package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request,X interface{}){
	// Read the body of the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// If there was an error reading the body, handle it
		return
	}

	// Unmarshal the body into the provided 'x' variable (interface{})
	if err := json.Unmarshal(body, X); err != nil {
		// If there was an error unmarshalling the JSON, handle it
		return
	}
}