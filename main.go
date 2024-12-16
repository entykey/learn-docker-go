package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define the index route
	r.GET("/", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("Hello, this is Go Gin version %s", gin.Version))
	})

	// Start the HTTP server on port 8080
	// r.Run(":8080")
	r.Run("0.0.0.0:8080")
}
