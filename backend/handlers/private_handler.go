package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/auth"
)

func (h handler) Me(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	claims, err := auth.GetTokenClaims(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": claims.Username})
}

func (h handler) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
