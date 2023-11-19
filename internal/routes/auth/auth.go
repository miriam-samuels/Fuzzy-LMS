package auth

import (
	"github.com/miriam-samuels/loan-management-backend/internal/controller/auth"
	"github.com/opensaucerer/barf"
)

func RegisterAuthRoutes(r *barf.SubRoute) {
	router := r.RetroFrame("/auth")

	router.Post("/signup", auth.UserSignUp)
	router.Post("/signin", auth.UserSignIn)
}
