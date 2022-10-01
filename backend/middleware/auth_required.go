package middleware

import (
	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/auth"
)

func JWTAuth(c *gin.Context) {
	// tokenString := c.GetHeader("Authorization")
	tokenString, err := c.Cookie("Authorization")
	if tokenString == "" || err != nil {
		c.JSON(401, gin.H{"error": "request does not contain an access token"})
		c.Abort()
		return
	}

	err = auth.ValidateToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Next()
}
