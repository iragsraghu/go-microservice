package api

import (
	"github/iragsraghu/go-microservice/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// User routes
	router.POST("/users", controllers.CreateUser)     // Create a new user
	router.GET("/users", controllers.GetUsers)        // Get all users
	router.GET("/users/:id", controllers.GetUserByID) // Get all users

	return router
}
