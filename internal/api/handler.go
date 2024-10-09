package api

import (
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(200, gin.H{
		"message": "User found",
		"userID":  userID,
	})
}
