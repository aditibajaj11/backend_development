package main

import (
	"log"
	"net/http"
	"os"
	"zocket_assignment/api"
	"zocket_assignment/db"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables for DB connection
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Initialize DB connection
	dbConn, err := db.Connect(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}

	// Set up the router and routes with DB connection
	router := api.SetRouter(dbConn)

	// Start the server on port 8080
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
