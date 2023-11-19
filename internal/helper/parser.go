package helper

import (
	"mime/multipart"
	"net/http"

	"github.com/opensaucerer/barf"
)

// This function takes in an interface for the request and returns a function which takes in a handler function and returns a handler

func ParseRequestBody(w http.ResponseWriter, r *http.Request, i interface{}) {
	err := barf.Request(r).Body().Format(i)
	if err != nil {
		SendJSONResponse(w, http.StatusBadRequest, false, "error parsing body:"+err.Error(), nil)
		return
	}
}

func ParseMultipartRequestBody(w http.ResponseWriter, r *http.Request) (multipart.File, error) {
	//  PARSE formdata including uploaded file
	head := barf.Request(r).Form().File().Get("file")
	f, err := head.Open()
	return f, err
}

