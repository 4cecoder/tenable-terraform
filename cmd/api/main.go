// Package api provides endpoints for interacting with Tenable's API to enhance services for company clients.
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// @title Tenable Client API
// @version 1.0
// @description The Tenable Client API enriches services provided to company clients by leveraging Tenable's API. It offers endpoints for managing assets, vulnerabilities, scans, and other Tenable-related functionalities to deliver robust and secure solutions.
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Define your API routes here
	// Example:
	// r.GET("/resources", getResources)
	// r.POST("/resources", createResource)
	// r.PUT("/resources/:id", updateResource)
	// r.DELETE("/resources/:id", deleteResource)

	// Run the server
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
