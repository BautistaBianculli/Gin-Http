package handlers

import (
	"GORUTINE/models"
	"GORUTINE/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	if user.Username != "BBIAN" || user.Password != "12345" {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "Invalid user or password",
		})
		return
	}

	token, err := token.GenerateToken(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
