package main

import (
	"log"
	"os"
	"pyvsgo/go/rest_api/config"
	"pyvsgo/go/rest_api/models"
	"pyvsgo/go/rest_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}()

	if err := db.AutoMigrate(&models.Book{}).Error; err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	bookRepo := models.NewBookRepository(db)
	router := gin.Default()
	router.RedirectTrailingSlash = true
	routes.BookRoutes(router, bookRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认端口
	}
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
