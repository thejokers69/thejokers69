package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thejokers69/thejokers69/internal/handlers"
)

func main() {
	// Set Gin mode
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	r := gin.Default()

	// Serve static files
	r.GET("/assets/*filepath", handlers.StaticHandler)
	r.GET("/thejokers69/assets/*filepath", handlers.StaticHandler)

	// Homepage
	r.GET("/", handlers.PageHandler)
	r.GET("/thejokers69/", handlers.PageHandler)

	// Page routes
	r.GET("/profile/", handlers.PageHandler)
	r.GET("/about-me/", handlers.PageHandler)
	r.GET("/skills/", handlers.PageHandler)
	r.GET("/projects/", handlers.PageHandler)
	r.GET("/contact/", handlers.PageHandler)

	// With baseurl prefix
	r.GET("/thejokers69/profile/", handlers.PageHandler)
	r.GET("/thejokers69/about-me/", handlers.PageHandler)
	r.GET("/thejokers69/skills/", handlers.PageHandler)
	r.GET("/thejokers69/projects/", handlers.PageHandler)
	r.GET("/thejokers69/contact/", handlers.PageHandler)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
