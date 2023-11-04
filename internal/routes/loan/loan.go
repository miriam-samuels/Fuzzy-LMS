package loan

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/controllers/loan"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
)

func RegisterLoanRoutes(r *mux.Router) {
	router := r.PathPrefix("/loan").Subrouter()

	// call middleware that validate auth
	router.Handle("/loans", middleware.ValidateAuth(loan.GetLoans)).Methods("GET")
	router.Handle("/create", middleware.ValidateAuth(loan.CreateLoanApplication)).Methods("POST")
}

//TODO: parse requests takes in the interface of the expected request body and the next http.handler that the returned function requires as parameter
