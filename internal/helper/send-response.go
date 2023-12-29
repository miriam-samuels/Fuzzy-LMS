package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

func SendResponse(w http.ResponseWriter, statusCode int, status bool, message string, data map[string]interface{}, err ...error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	e := json.NewEncoder(w).Encode(types.Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
	if e != nil {
		// Handle the error (e.g., log it)
		fmt.Printf("Error:: %v", e.Error())
	}
	if status == false && len(err) > 0 {
		log.Printf(message+":: %v", err[0].Error())
	}
}
