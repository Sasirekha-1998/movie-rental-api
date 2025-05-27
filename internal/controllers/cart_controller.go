package controllers

import (
	"net/http"
	"strconv"

	"movie-rental-api/internal/services"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	CartService *services.CartService
}

func NewCartController(cartService *services.CartService) *CartController {
	return &CartController{CartService: cartService}
}

func (cc *CartController) AddToCart(c *gin.Context) {
	userIDStr := c.Query("user_id")
	movieIDStr := c.Query("movie_id")

	userID, err1 := strconv.Atoi(userIDStr)
	movieID, err2 := strconv.Atoi(movieIDStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id or movie_id"})
		return
	}

	err := cc.CartService.AddMovieToCart(userID, movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie added to cart"})
}

func (cc *CartController) GetCart(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	cartItems, err := cc.CartService.GetUserCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}
