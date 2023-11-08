package helper

import (
	"net/http"

	"github.com/opensaucerer/barf"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, status bool, message string, data map[string]interface{}) {
	barf.Response(w).Status(statusCode).JSON(barf.Res{
		Status:  status,
		Data:    data,
		Message: message,
	})
}
