package helper

import (
	"encoding/json"
	"net/http"
)

// This function takes in an interface for the request and returns a function which takes in a handler function and returns a handler

func ParseRequestBody(w http.ResponseWriter, r *http.Request, i interface{}) error {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		if err.Error() != "EOF" {
			// send response on unable to parse body
			SendJSONResponse(w, http.StatusBadRequest, false, "error parsing body:"+err.Error(), nil)
			return err
		}
	}
	return nil
}
