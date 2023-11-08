package media

import (
	"github.com/miriam-samuels/loan-management-backend/internal/controllers/media"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
	"github.com/opensaucerer/barf"
)

func RegisterMediaRoutes(r *barf.SubRoute) {
	router := r.RetroFrame("/media")

	router.Post("/upload",media.UploadMedia, middleware.ValidateAuth)
}
