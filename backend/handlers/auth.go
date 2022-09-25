package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/global"
	"mikesprogram.com/tenbeat/models"
)

func (h handler) Login(c *gin.Context) {
	session := sessions.Default(c)

	// type user struct {
	// 	username string
	// 	password string
	// }

	var u models.User

	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println(u)

	// Validate form input
	if strings.Trim(u.Username, " ") == "" || strings.Trim(u.Password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if u.Username != "hello" || u.Password != "itsme" {
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
	session.Set(global.Userkey, u.Username) // In real world usage you'd set this to the users ID
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	fmt.Println(c.Cookie("mysession"))

	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func (h handler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(global.Userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(global.Userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	fmt.Println(c.Cookie("mysession"))
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
