package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/miriam-samuels/loan-management-backend/internal/helper"
)

// function to validate user authorization
func ValidateAuth(nextHandler http.Handler) http.Handler {
	// returns a handler function which calls the next handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// TODO: send respose that user is not authorized
			return
		}

		// split the header e.g Bearer somerandomtoken
		authHeaderSplit := strings.Split(authHeader, " ")

		// Check if auth header has a lenght of 2 after spilting
		if len(authHeaderSplit) != 2 {
			// TODO: send response on invalid auth header
			return
		}

		// check if auth type is bearer
		if authHeaderSplit[0] != "Bearer" {
			//	TODO: invalid auth type
			return
		}

		// retieve token from auth header
		token := authHeaderSplit[1]
		if token == "" {
			// TODO: send response on user not logged in
			return
		}

		// validation logic for token (convert _ to claims)
		_, valid := helper.VerifyJWT(token)
		if !valid {
			// TODO: send response on invalid token provided
			return
		}

		// TODO: Verify user using token claims

		// call nect handler
		nextHandler.ServeHTTP(w, r)
	})
}

// This function takes in an interface for the request and returns a function which takes in a handler function and returns a handler
func ParseRequest(i interface{}) func(http.HandlerFunc) http.Handler {
	return func(nextHandler http.HandlerFunc) http.Handler {
		// anonymous function using handlerfunc returns a handler
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := json.NewDecoder(r.Body).Decode(i)
			if err != nil {
				if err.Error() != "EOF" {
					// TODO: send response on unable to parse body
					return
				}

				nextHandler(w, r)
			}
		})
	}
}
