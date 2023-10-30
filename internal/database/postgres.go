package database

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

const DRIVER_NAME = "postgres"

// function to create new postgres client to interface with db
func NewPostgresClient(uri string) (client *sql.DB, connectionError error) {
	client, connectionError = sql.Open(DRIVER_NAME, uri)
	client.SetConnMaxLifetime(time.Second * 30)
	fmt.Println("connected to database")
	return
}
