package repositories

import (
	"database/sql"
	"movie-rental-api/internal/models"
)

type CartRepository struct {
	DB *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) AddToCart(userID, movieID int) error {
	_, err := r.DB.Exec(`
        INSERT INTO cart (user_id, movie_id)
        VALUES ($1, $2)
        ON CONFLICT (user_id, movie_id) DO NOTHING
    `, userID, movieID)
	return err
}

func (r *CartRepository) GetCartItems(userID int) ([]models.CartItem, error) {
	rows, err := r.DB.Query(`
        SELECT id, user_id, movie_id, added_at
        FROM cart
        WHERE user_id = $1
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.CartItem
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ID, &item.UserID, &item.MovieID, &item.AddedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
