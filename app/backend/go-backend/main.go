package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default gin router
	r := gin.Default()

	// Define your routes
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Go backend with Gin!")
	})

	// Add more routes as needed
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Run the server on port 8080
	r.Run(":8080") // equivalent to http.ListenAndServe(":8080", r)
}
