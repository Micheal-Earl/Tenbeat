package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mikesprogram.com/tenbeat/models"
)

func (h handler) GetAllBooks(c *gin.Context) {
	var books []models.Book

	result := h.DB.Find(&books)
	if result.Error != nil {
		http.Error(c.Writer, "DB find failed", http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, books)
}

func (h handler) GetBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(c.Writer, "Invalid ID", http.StatusBadRequest)
	}

	var book models.Book

	result := h.DB.First(&book, id)
	if result.Error != nil {
		http.Error(c.Writer, "DB find failed", http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, book)
}

func (h handler) AddBook(c *gin.Context) {
	var book models.Book

	err := c.BindJSON(&book)
	if err != nil {
		http.Error(c.Writer, "Cannot bind JSON", http.StatusBadRequest)
	}

	result := h.DB.Create(&book)
	if result.Error != nil {
		http.Error(c.Writer, "DB insert failed", http.StatusInternalServerError)
	}

	c.Status(http.StatusCreated)
}

func (h handler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(c.Writer, "Invalid ID", http.StatusBadRequest)
	}

	var updatedBook models.Book

	err = c.BindJSON(&updatedBook)
	if err != nil {
		http.Error(c.Writer, "Cannot bind JSON", http.StatusBadRequest)
	}

	var book models.Book

	result := h.DB.First(&book, id)
	if result.Error != nil {
		http.Error(c.Writer, "DB insert failed", http.StatusInternalServerError)
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc

	h.DB.Save(&book)

	c.Status(http.StatusOK)
}

func (h handler) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(c.Writer, "Invalid ID", http.StatusBadRequest)
	}

	var book models.Book

	result := h.DB.First(&book, id)
	if result.Error != nil {
		http.Error(c.Writer, "DB find failed", http.StatusInternalServerError)
	}

	h.DB.Delete(&book)

	c.Status(http.StatusOK)
}
