package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/db"
	"mikesprogram.com/tenbeat/global"
	"mikesprogram.com/tenbeat/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	router := gin.Default()

	// Middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	router.Use(sessions.Sessions("mysession", sessions.NewCookieStore(global.Secret)))

	// Routes
	router.GET("/books", h.GetAllBooks)
	router.GET("/books/:id", h.GetBook)
	router.POST("/books", h.AddBook)
	router.PUT("/books/:id", h.UpdateBook)
	router.DELETE("/books/:id", h.DeleteBook)

	router.POST("/login", h.Login)
	router.GET("/logout", h.Logout)

	// Private group, require authentication to access
	private := router.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
	}

	err := router.Run("localhost:9090")
	if err != nil {
		return
	}
}

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

func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(global.Userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
