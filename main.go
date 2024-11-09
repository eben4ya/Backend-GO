package main

import (
	"log"
	"net/http"

	"book_store/config"
	"book_store/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Connect to MongoDB
    config.ConnectDB()

    // Initialize router
    r := mux.NewRouter()

    // Register routes
    routes.RegisterRoutes(r)

    // Start the server
    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}
