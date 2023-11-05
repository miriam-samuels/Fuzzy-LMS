package user

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/controllers/user"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
)

func RegisterUserRoutes(r *mux.Router) {
	router := r.PathPrefix("/user").Subrouter()

	router.Handle("/profile/me", middleware.ValidateAuth(user.GetProfile)).Methods("GET")
	router.Handle("/profile", middleware.ValidateAuth(user.UpdateProfile)).Methods("PATCH")
}
