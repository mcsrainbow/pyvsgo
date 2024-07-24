package main

import (
	"pyvsgo/go/rest_api_v1/book"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&book.Book{})

	// Initialize the book package with the database
	book.Init(db)

	r := gin.Default()

	// API routes
	r.GET("/api/show_books", book.ShowBooks)
	r.POST("/api/add_book", book.AddBook)
	r.DELETE("/api/delete_book/:id", book.DeleteBook)
	r.PUT("/api/update_book/:id", book.UpdateBook)

	r.Run(":8080") // Run on port 8080
}
