package main

import (
	"log"
	"movie-rental-api/config"
	"movie-rental-api/internal/controllers"
	"movie-rental-api/internal/database"
	"movie-rental-api/internal/repositories"
	"movie-rental-api/internal/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type MovieController struct {
	Service services.MovieService
}

func main() {
	// Connect to the database
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	port := cfg.Server.Port

	router := gin.Default()

	//story1
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello , Welcome to the Movie Rental Dashboard!",
		})
	})

	//story2
	router.GET("/movies", controllers.GetAllMovies)

	//story3
	movieRepo := repositories.NewMovieRepository(db)
	movieService := services.NewMovieService(movieRepo)
	movieController := controllers.NewMovieController(movieService)
	router.GET("/movies/filter", movieController.GetFilteredMovies)

	log.Println("✅ Server running on :" + port)
	router.Run(":" + port) // Use the port from the config
}
