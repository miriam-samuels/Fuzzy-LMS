package helper

import (
	"log"
	"net/http"

	"github.com/opensaucerer/barf"
)

func SendResponse(w http.ResponseWriter, statusCode int, status bool, message string, data map[string]interface{}, err ...error) {
	barf.Response(w).Status(statusCode).JSON(barf.Res{
		Status:  status,
		Data:    data,
		Message: message,
	})

	if status == false && len(err) > 0 {
		log.Printf(message+":: %v", err[0].Error())
	}
}
