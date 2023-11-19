package main

import (
	"log"
	"github.com/miriam-samuels/loan-management-backend/internal/constants"
	"github.com/miriam-samuels/loan-management-backend/internal/database"
	"github.com/miriam-samuels/loan-management-backend/internal/storage"
	v1 "github.com/miriam-samuels/loan-management-backend/internal/version/v1"
	"github.com/opensaucerer/barf"
)

// connection port and host for local environment
const (
	CONN_PORT = "6000"
)

func init() {
	// Find .env file
	if err := barf.Env(constants.Env, ".env"); err != nil {
		barf.Logger().Fatal(err.Error())
	}

	// Connect Database
	client, err := database.NewPostgresClient(constants.Env.LoanDbDatasourceUri)
	if err != nil {
		log.Fatal("error connecting to database :: ", err)
	}

	// set loandb to created client
	database.LoanDb = client

	// Connect to Storage Bucket
	bucket, err := storage.NewFirebaseBucket("serviceAccountKey.json", constants.Env.StorageBucket)
	if err != nil {
		log.Fatal("error getting storage bucket ::", err)
	}
	// set loan bucket to gotten bucket
	storage.LoanBucket = bucket
}

func main() {
	// Get port if it exists in env file
	port := constants.Env.ConnectionPort
	// check if port exists in env file else use constant
	if port == "" {
		port = CONN_PORT
	}

	// register routes with versioning
	v1.Routes()


	//  Defer connection to db close
	defer database.LoanDb.Close()

	if err := barf.Stark(barf.Augment{
		Host: "",
		Port:         port,
		Logging:      barf.Allow(), // enable request logging
		Recovery:     barf.Allow(), // enable panic recovery so barf returns a 500 error instead of crashing
		ReadTimeout:  30,
		WriteTimeout: 30,
		CORS: &barf.CORS{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			AllowedHeaders:   []string{"Authorization", "Content-Type"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		},
	}); err != nil {
		barf.Logger().Fatal(err.Error())
	}

	// create & start server
	if err := barf.Beck(); err != nil {
		// barf exposes a logger instance
		barf.Logger().Fatal(err.Error())
	}
}
