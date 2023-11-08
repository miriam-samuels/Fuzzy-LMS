package user

import (
	"github.com/miriam-samuels/loan-management-backend/internal/controllers/user"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
	"github.com/opensaucerer/barf"
)

func RegisterUserRoutes(r *barf.SubRoute) {
	router := r.RetroFrame("/user")

	router.Get("/profile/me", user.GetProfile, middleware.ValidateAuth)
	router.Post("/profile", user.UpdateProfile, middleware.ValidateAuth)
}
