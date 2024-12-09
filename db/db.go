package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Connect initializes a connection to the database using the provided configuration
func Connect(dbHost, dbPort, dbUser, dbPassword, dbName string) (*sql.DB, error) {
	// Create the connection string using the provided parameters
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open the connection to the database
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
		return nil, err
	}

	// Test the connection
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return dbConn, nil
}
