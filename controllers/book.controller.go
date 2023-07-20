package controllers

import (
	"myGin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post Book
func (db *DBController) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := db.Database.Create(&book)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create book"})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"results": book})
		return
	}
}

// Get book by id
func (db *DBController) GetBookById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	tx := db.Database.First(&book, id)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"results": book})
		return
	}
}

// Get book list
func (db *DBController) GetBookLists(c *gin.Context) {
	var books []models.Book
	tx := db.Database.Find(&books)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"results": books})
		return
	}
}

// Update book
func (db *DBController) UpdateBook(c *gin.Context) {
	var book models.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	tx := db.Database.Updates(&book)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot update book"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"results": "book updated"})
		return
	}
}

// delete book
func (db *DBController) DeleteBookById(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	tx := db.Database.Delete(&book, id)
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete book"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"results": "book deleted"})
		return
	}
}
