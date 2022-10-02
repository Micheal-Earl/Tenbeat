package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/auth"
	"mikesprogram.com/tenbeat/models"
)

// TODO: Reduce code reuse by breaking error handling into its
// own function + authorization check into its own function

func (h handler) GetAllPosts(c *gin.Context) {
	var posts []models.Post

	result := h.DB.Preload("Owner").Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		c.Abort()
		return
	}

	sanitizedPosts := make([]models.SanitizedPost, len(posts))
	for i, post := range posts {
		sanitizedPosts[i] = post.SanitizePost()
	}

	c.JSON(http.StatusOK, sanitizedPosts)
}

func (h handler) GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var post models.Post

	result := h.DB.Preload("Owner").First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		c.Abort()
		return
	}

	sanitizedPost := post.SanitizePost()

	c.JSON(http.StatusOK, sanitizedPost)
}

func (h handler) CreatePost(c *gin.Context) {
	var post models.Post

	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

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

	post.OwnerID = claims.ID

	result := h.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Post created"})
}

func (h handler) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var updatedPost models.Post

	err = c.BindJSON(&updatedPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var post models.Post

	result := h.DB.Preload("Owner").First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		c.Abort()
		return
	}

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

	if claims.ID != post.Owner.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized Edit"})
		c.Abort()
		return
	}

	post.Title = updatedPost.Title
	post.Content = updatedPost.Content

	h.DB.Save(&post)

	c.Status(http.StatusOK)
}

func (h handler) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	var post models.Post

	result := h.DB.Preload("Owner").First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		c.Abort()
		return
	}

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

	fmt.Println(claims.ID, " ", post.Owner.ID)

	if claims.ID != post.Owner.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized Delete"})
		c.Abort()
		return
	}

	h.DB.Delete(&post)

	c.Status(http.StatusOK)
}
