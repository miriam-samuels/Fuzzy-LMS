package borrower

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/controller/borrower"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
)

func RegisterBorrowerRoutes(r *mux.Router) {
	router := r.PathPrefix("/borrower").Subrouter()

	router.Handle("/all", middleware.ValidateAuth(borrower.GetBorrowers)).Methods("GET")
	router.Handle("/{id}", middleware.ValidateAuth(borrower.GetBorrowerById)).Methods("GET")
}
