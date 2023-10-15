package v1

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/auth"
)

func Routes(router *mux.Router) {
	// handle versioning
	r := router.PathPrefix("/v1").Subrouter()

	// Register routes
	auth.RegisterAuthRoutes(r)

}
