package auth

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/controllers/auth"
)

func RegisterAuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	// call middleware that verify then another middleware which parses
	// parse requests takes in the interface of the expected request body and the next http.handler that the returned function requires as parameter
	router.HandleFunc("/signup", auth.LenderSignUp).Methods("POST")
	router.HandleFunc("/signin", auth.LenderSignIn).Methods("POST")

}
