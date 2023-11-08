package v1

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/auth"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/loan"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/media"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/user"
)

func Routes(router *mux.Router) {
	// handle versioning
	r := router.PathPrefix("/v1").Subrouter()

	// Register routes
	auth.RegisterAuthRoutes(r)
	loan.RegisterLoanRoutes(r)
	user.RegisterUserRoutes(r)
	media.RegisterMediaRoutes(r)
}