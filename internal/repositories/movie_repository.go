package repositories

import (
	"database/sql"
	"fmt"
	"movie-rental-api/internal/database"
	"movie-rental-api/internal/models"
)

func GetAllMovies() ([]models.Movie, error) {
	rows, err := database.DB.Query("SELECT id, title, year, imdb_id, type, poster, description, genre, language, rating, duration ,actors FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		err := rows.Scan(&m.ID, &m.Title, &m.Year, &m.ImdbID, &m.Type, &m.Poster, &m.Description, &m.Genre, &m.Language, &m.Rating, &m.Duration, &m.Actors)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}

// repositories/movie_repository.go
type MovieRepository struct {
	DB *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (r *MovieRepository) GetMoviesByFilters(filters map[string]string) ([]models.Movie, error) {
	baseQuery := `SELECT id, title, year, imdb_id, type, poster, description, genre, language, rating, duration, COALESCE(actors, '') FROM movies WHERE 1=1`
	args := []interface{}{}
	i := 1

	for key, value := range filters {
		// Prevent SQL injection by allowing only known safe columns
		switch key {
		case "title", "genre", "type", "language", "rating", "actors":
			baseQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, i)
			args = append(args, "%"+value+"%")
			i++
		case "year":
			baseQuery += fmt.Sprintf(" AND %s = $%d", key, i)
			args = append(args, value)
			i++
		default:
			// Ignore unknown keys
		}
	}

	rows, err := r.DB.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		err := rows.Scan(
			&m.ID, &m.Title, &m.Year, &m.ImdbID, &m.Type, &m.Poster,
			&m.Description, &m.Genre, &m.Language, &m.Rating, &m.Duration, &m.Actors,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}
	return movies, nil
}
