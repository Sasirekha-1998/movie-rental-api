package controllers

import (
	"database/sql"
	"errors"
	"movie-rental-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(c *gin.Context) {
	services.GetAllMovies(c)
}

type MovieController struct {
	Service *services.MovieService
}

func NewMovieController(service *services.MovieService) *MovieController {
	return &MovieController{Service: service}
}

func (mc *MovieController) GetFilteredMovies(c *gin.Context) {
	// Get all query parameters as a map
	filters := map[string]string{}
	for key, values := range c.Request.URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	movies, err := mc.Service.GetMoviesByFilters(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (mc *MovieController) GetMovieByID(c *gin.Context) {
	id := c.Query("id")
	movie, err := mc.Service.GetMovieByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie details"})
		}
		return
	}
	c.JSON(http.StatusOK, movie)
}
