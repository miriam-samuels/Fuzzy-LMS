package v1

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/auth"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/borrower"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/dashboard"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/loan"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/media"
	"github.com/miriam-samuels/loan-management-backend/internal/routes/profile"
)

func Routes(router *mux.Router) {
	// handle versioning
	r := router.PathPrefix("/v1").Subrouter()

	// Register routes
	auth.RegisterAuthRoutes(r)
	loan.RegisterLoanRoutes(r)
	profile.RegisterProfileRoutes(r)
	media.RegisterMediaRoutes(r)
	dashboard.RegisterDashboardRoutes(r)
	borrower.RegisterBorrowerRoutes(r)
}
