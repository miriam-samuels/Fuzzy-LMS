package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/miriam-samuels/loan-management-backend/internal/database"
	"github.com/miriam-samuels/loan-management-backend/internal/storage"
	v1 "github.com/miriam-samuels/loan-management-backend/internal/version/v1"
	"github.com/rs/cors"
)

// connection port and host for local environment
const (
	CONN_PORT = "6000"
)

func init() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Connect Database
	client, err := database.NewPostgresClient(os.Getenv("LOAN_DB_DATASOURCE_URI"))
	if err != nil {
		log.Fatal("error connecting to database :: ", err)
	}

	// set loandb to created client
	database.LoanDb = client

	// Connect to Storage Bucket
	bucket, err := storage.NewFirebaseBucket("serviceAccountKey.json", os.Getenv("STORAGE_BUCKET"))
	if err != nil {
		log.Fatal("error getting storage bucket ::", err)
	}
	// set loan bucket to gotten bucket
	storage.LoanBucket = bucket
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

	//  cross origin
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		// Debug:            true,
	}).Handler(router)

	// add more configurations to server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	// start server
	fmt.Println("starting server on port :: " + port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
