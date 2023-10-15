package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/miriam-samuels/loan-management-backend/internal/database"
	v1 "github.com/miriam-samuels/loan-management-backend/internal/version/v1"
)

// connection port and host for local environment
const (
	CONN_PORT = "3000"
)

func init() {
	// Connect Database
	client, err := database.NewPostgresClient(os.Getenv("LOAN_DB_DATASOURCE_URI"))
	if err != nil {
		log.Fatal("error connecting to database :: ", err)
	}

	// set loandb to created client
	database.LoanDb = client
}

func main() {
	// Get port if it exists in env file
	port := os.Getenv("PORT")
	// check if port exists in env file else use constant
	if port == "" {
		port = CONN_PORT
	}

	// create new router
	router := mux.NewRouter().StrictSlash(true)

	// register routes with versioning
	v1.Routes(router)

	//  Defer connection to db close
	defer database.LoanDb.Close()

	// add more configurations to server
	server := http.Server{
		Addr:        ":" + port,
		Handler:     router,
		ReadTimeout: time.Second * 90,
		WriteTimeout: time.Second * 90,
	}
	// listen on port
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
