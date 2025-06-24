// Entry point of the application
package main

import (
	"campus-compass/backend/auth"
	"campus-compass/backend/config"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"campus-compass/backend/profile"
)

func main() {
	// Initialize configuration
	config.InitDB()
	defer config.DB.Close()

	// Setup Gin
	r := gin.Default()

	// CORS Middleware
	// r.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(204)
	// 		return
	// 	}
		
	// 	c.Next()
	// })
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Your Next.js URL
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup routes
	auth.SetupRoutes(r)
	// Setup routes
	profile.SetupRoutes(r) // Add this line

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on :%s", port)
	log.Fatal(r.Run(":" + port))
}