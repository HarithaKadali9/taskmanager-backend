package main

import (
	"taskmanager/database"
	"taskmanager/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.Connect()

	// Initialize Gin router
	router := gin.Default()

	// Apply default CORS middleware
	router.Use(cors.Default())

	// Register all routes
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
