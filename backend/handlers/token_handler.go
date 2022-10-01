package handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"mikesprogram.com/tenbeat/auth"
// 	"mikesprogram.com/tenbeat/models"
// )

// type TokenRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func (h handler) GenerateToken(c *gin.Context) {
// 	var request TokenRequest
// 	var user models.User

// 	err := c.ShouldBindJSON(&request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	// check if email exists and password is correct
// 	record := h.DB.Where("email = ?", request.Email).First(&user)
// 	if record.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
// 		c.Abort()
// 		return
// 	}

// 	credentialError := user.CheckPassword(request.Password)
// 	if credentialError != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
// 		c.Abort()
// 		return
// 	}

// 	tokenString, err := auth.GenerateJWT(user.Email, user.Username)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	c.SetCookie("Authorization", tokenString, 2147383647, "/", "localhost:9090", true, true)

// 	c.JSON(http.StatusOK, gin.H{"token": tokenString})
// }
