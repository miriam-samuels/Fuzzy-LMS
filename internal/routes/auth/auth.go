package auth

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/controllers/auth"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

func RegisterAuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	// call middleware that verify then another middleware which parses 
	// parse requests takes in the interface of the expected request body and the next http.handler that the returned function requires as parameter
	router.Handle("/signup", middleware.ValidateAuth(middleware.ParseRequest(&types.SignUpCred{})(auth.UserSignUp))).Methods("POST")
	router.Handle("/signin", middleware.ValidateAuth(middleware.ParseRequest(&types.SignInCred{})(auth.UserSignIn))).Methods("POST")

}
