package controllers

import (
	"github/iragsraghu/go-microservice/internal/models"
	"github/iragsraghu/go-microservice/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := services.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers handles fetching all users from the database
func GetUsers(c *gin.Context) {
	var users []models.User
	if result := services.DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID handles fetching a user by their ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id") // Get the ID from the URL parameter
	var user models.User

	if result := services.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
