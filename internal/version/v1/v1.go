package v1

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/auth"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/loan"
)

func Routes(router *mux.Router) {
	// handle versioning
	r := router.PathPrefix("/v1").Subrouter()

	// Register routes
	auth.RegisterAuthRoutes(r)
	loan.RegisterLoanRoutes(r)

}