package profile

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/controller/profile"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
)

func RegisterProfileRoutes(r *mux.Router) {
	router := r.PathPrefix("/profile").Subrouter()

	router.Handle("/me", middleware.ValidateAuth(profile.GetProfile)).Methods("GET")
	router.Handle("/me", middleware.ValidateAuth(profile.UpdateProfile)).Methods("PATCH")
}
