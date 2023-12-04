package dashboard

import (
	"github.com/miriam-samuels/loan-management-backend/internal/controller/dashboard"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
	"github.com/opensaucerer/barf"
)

func RegisterDashboardRoutes(r *barf.SubRoute) {
	router := r.RetroFrame("/dashboard")

	router.Get("/loan-stats", dashboard.GetDashboardData, middleware.ValidateAuth)
}
