package services

import (
	"movie-rental-api/internal/models"
	"movie-rental-api/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieService struct {
	Repo *repositories.MovieRepository
}

func NewMovieService(repo *repositories.MovieRepository) *MovieService {
	return &MovieService{Repo: repo}
}

func GetAllMovies(c *gin.Context) {
	movies, err := repositories.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// func (s *MovieService) GetFilteredMovies(genre, actor string, year int) ([]models.Movie, error) {
// 	return s.Repo.GetFilteredMovies(genre, actor, year)
// }

func (s *MovieService) GetMoviesByFilters(filters map[string]string) ([]models.Movie, error) {
	return s.Repo.GetMoviesByFilters(filters)
}

func (s *MovieService) GetMovieByID(id string) (*models.Movie, error) {
	return s.Repo.GetMovieByID(id)
}
