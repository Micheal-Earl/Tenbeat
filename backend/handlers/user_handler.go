package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/models"
)

func (h handler) RegisterUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	err = user.HashPassword(user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	result := h.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{"userId": user.ID, "email": user.Email, "username": user.Username},
	)
}

func (h handler) ValidatePassword(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var userCompare models.User

	result := h.DB.First(&userCompare, "username = ?", user.Username)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		c.Abort()
		return
	}

	err = userCompare.CheckPassword(user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"valid": false})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"valid": true})
}

// func (h handler) Login(c *gin.Context) {
// 	session := sessions.Default(c)

// 	var user models.User

// 	err := c.BindJSON(&user)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
// 		return
// 	}

// 	// Validate form input
// 	if strings.Trim(user.Username, " ") == "" || strings.Trim(user.PasswordHash, " ") == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
// 		return
// 	}

// 	// Check for username and password match, usually from a database
// 	userFromDB, err := h.GetUser(c, user.Username)
// 	if err != nil {
// 		http.Error(c.Writer, "Could not find user", http.StatusInternalServerError)
// 	}

// 	if user.PasswordHash != userFromDB.PasswordHash {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
// 		return
// 	}

// 	// Save the username in the session
// 	// sessionConfig := sessions.Options{
// 	// 	Path: "/",
// 	// 	//Domain: "http://127.0.0.1",
// 	// 	MaxAge: 2592000,
// 	// 	Secure: true,
// 	// }
// 	// session.Options(sessionConfig)
// 	//session.Set(userFromDB.Id, userFromDB.Username) // In real world usage you'd set this to the users ID
// 	session.Set(global.Userkey, userFromDB.Username)
// 	err = session.Save()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
// 		return
// 	}

// 	//fmt.Println(c.Cookie("mysession"))

// 	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
// }

// func (h handler) Logout(c *gin.Context) {
// 	session := sessions.Default(c)

// 	sessionToken := session.Get(global.Userkey)
// 	if sessionToken == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
// 		return
// 	}

// 	session.Delete(global.Userkey)
// 	if err := session.Save(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
// }

// func (h handler) GetUser(c *gin.Context, username string) (*models.User, error) {
// 	var user models.User

// 	result := h.DB.Table("users").Where("username = ?", username).First(&user)

// 	return &user, result.Error
// }
