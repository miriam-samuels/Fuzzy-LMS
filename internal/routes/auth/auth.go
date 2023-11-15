package auth

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/controllers/auth"
)

func RegisterAuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/signup", auth.UserSignUp).Methods("POST")
	router.HandleFunc("/signin", auth.UserSignIn).Methods("POST")
}
 