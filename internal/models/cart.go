package models

import "time"

type CartItem struct {
	ID      int       `json:"id"`
	UserID  int       `json:"user_id"`
	MovieID int       `json:"movie_id"`
	AddedAt time.Time `json:"added_at"`
}
