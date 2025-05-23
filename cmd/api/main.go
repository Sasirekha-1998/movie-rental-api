package main

import (
	"fmt"
	"log"
	"movie-rental-api/config"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the Gin router
	router := gin.Default()

	// Define a simple route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	port := cfg.Server.Port
	fmt.Println("Port:", port)

	// Start the server on the configured port
	router.Run(":" + port) // config.Server.Port
}
