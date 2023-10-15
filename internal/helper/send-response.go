package helper

import (
	"encoding/json"
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, status bool, message string, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(types.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}
