package services

import (
	"movie-rental-api/internal/models"
	"movie-rental-api/internal/repositories"
)

type CartService struct {
	CartRepo *repositories.CartRepository
}

func NewCartService(cartRepo *repositories.CartRepository) *CartService {
	return &CartService{CartRepo: cartRepo}
}

func (s *CartService) AddMovieToCart(userID, movieID int) error {
	return s.CartRepo.AddToCart(userID, movieID)
}

func (s *CartService) GetUserCart(userID int) ([]models.CartItem, error) {
	return s.CartRepo.GetCartItems(userID)
}
