package book

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Book struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	BookName string `json:"book_name"`
	AddTime  int64  `json:"add_time"`
}

var db *gorm.DB

// Init initializes the package with a database connection
func Init(database *gorm.DB) {
	db = database
}

func ShowBooks(c *gin.Context) {
	var books []Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBook.AddTime = time.Now().Unix()
	db.Create(&newBook)
	c.JSON(http.StatusCreated, newBook)
}

func DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book Book
	if db.First(&book, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusNoContent, nil)
}
