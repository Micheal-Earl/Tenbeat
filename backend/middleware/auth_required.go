package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/auth"
	"mikesprogram.com/tenbeat/global"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(global.Userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	fmt.Println(session)
	fmt.Println(c.Cookie("mysession"))

	// Continue down the chain to handler etc
	c.Next()
}

func JWTAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "request does not contain an access token"})
		c.Abort()
		return
	}

	err := auth.ValidateToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Next()
}

// func JWTAuth() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		tokenString := context.GetHeader("Authorization")
// 		if tokenString == "" {
// 			context.JSON(401, gin.H{"error": "request does not contain an access token"})
// 			context.Abort()
// 			return
// 		}
// 		err := auth.ValidateToken(tokenString)
// 		if err != nil {
// 			context.JSON(401, gin.H{"error": err.Error()})
// 			context.Abort()
// 			return
// 		}
// 		context.Next()
// 	}
// }
