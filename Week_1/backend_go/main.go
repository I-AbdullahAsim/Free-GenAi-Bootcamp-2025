package main

import (
	"github.com/gin-gonic/gin"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/api"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/db"
)

func main() {
	// Initialize database
	db.GetDB()

	// Create Gin router with default middleware
	router := gin.Default()

	// Setup API routes
	api.SetupRoutes(router)

	// Start server
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}