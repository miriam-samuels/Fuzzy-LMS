package media

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/controllers/media"
	"github.com/miriam-samuels/loan-management-backend/internal/middleware"
)

func RegisterMediaRoutes(r *mux.Router) {
	router := r.PathPrefix("/media").Subrouter()

	router.Handle("/upload", middleware.ValidateAuth(media.UploadMedia)).Methods("POST")
}
