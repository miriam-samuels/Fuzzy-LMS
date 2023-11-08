package v1

import (
	"github.com/miriam-samuels/loan-management-backend/internal/routes/auth"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/loan"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/media"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/user"
	"github.com/opensaucerer/barf"
)

func Routes() {
	// handle versioning
	r := barf.RetroFrame("/v1")

	// Register routes
	auth.RegisterAuthRoutes(r)
	loan.RegisterLoanRoutes(r)
	user.RegisterUserRoutes(r)
	media.RegisterMediaRoutes(r)
}
