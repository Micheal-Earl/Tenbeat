package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/db"
	"mikesprogram.com/tenbeat/global"
	"mikesprogram.com/tenbeat/handlers"
	"mikesprogram.com/tenbeat/middleware"
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	router.POST("/register", h.RegisterUser)
	router.POST("/token", h.GenerateToken)
	router.POST("/valid", h.ValidatePassword)

	// router.POST("/login", h.Login)
	// router.GET("/logout", h.Logout)

	// Private group, require authentication to access
	private := router.Group("/private")
	private.Use(middleware.JWTAuth)
	{
		private.GET("/me", h.Me)
		private.GET("/status", h.Status)
	}

	srv := &http.Server{
		Addr:    "localhost:9090",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
