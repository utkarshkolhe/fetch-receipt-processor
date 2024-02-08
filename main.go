package main

import (
	"log"
	"net/http"
	"utkarsh/Fetch/api/routes"

	"github.com/gin-gonic/gin"
)

// Entry point of the whole program
func main() {
	// Create a gin router to handle requests
	router := gin.Default()
	// Define all API routes
	routes.SetupRoutes(router)

	// Start the server on 8080. Port number can be changed.
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
