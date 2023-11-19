package loan

import (
	"github.com/miriam-samuels/loan-management-backend/internal/controller/loan"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
	"github.com/opensaucerer/barf"
)

func RegisterLoanRoutes(r *barf.SubRoute) {
	router := r.RetroFrame("/loan")

	// call middleware that validate auth
	router.Get("/loans", loan.GetLoans, middleware.ValidateAuth)
	router.Post("/create", loan.CreateLoanApplication, middleware.ValidateAuth)
}

//TODO: parse requests takes in the interface of the expected request body and the next http.handler that the returned function requires as parameter
