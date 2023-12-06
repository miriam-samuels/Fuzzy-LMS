package dashboard

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/controller/dashboard"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
)

func RegisterDashboardRoutes(r *mux.Router) {
	router := r.PathPrefix("/dashboard").Subrouter()

	router.Handle("/stats", middleware.ValidateAuth(dashboard.GetDashboardData)).Methods("GET")
}
