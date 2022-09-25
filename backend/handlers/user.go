package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/global"
	"mikesprogram.com/tenbeat/models"
)

func (h handler) Login(c *gin.Context) {
	session := sessions.Default(c)

	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Validate form input
	if strings.Trim(user.Username, " ") == "" || strings.Trim(user.Password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	userFromDB, err := h.GetUser(c, user.Username)
	if err != nil {
		http.Error(c.Writer, "Could not find user", http.StatusInternalServerError)
	}

	if user.Password != userFromDB.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	// sessionConfig := sessions.Options{
	// 	Path: "/",
	// 	//Domain: "http://127.0.0.1",
	// 	MaxAge: 2592000,
	// 	Secure: true,
	// }
	// session.Options(sessionConfig)
	//session.Set(userFromDB.Id, userFromDB.Username) // In real world usage you'd set this to the users ID
	session.Set(global.Userkey, userFromDB.Username)
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	//fmt.Println(c.Cookie("mysession"))

	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func (h handler) Logout(c *gin.Context) {
	session := sessions.Default(c)

	sessionToken := session.Get(global.Userkey)
	if sessionToken == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete(global.Userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func (h handler) GetUser(c *gin.Context, username string) (*models.User, error) {
	var user models.User

	result := h.DB.Table("users").Where("username = ?", username).First(&user)

	return &user, result.Error
}
