package routes

import (
	"pyvsgo/go/rest_api/controllers"
	"pyvsgo/go/rest_api/models"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine, repo *models.BookRepository) {
	bookController := controllers.NewBookController(repo)

	bookRoutes := router.Group("/api")
	{
		bookRoutes.GET("/books", bookController.GetBooks)
		bookRoutes.POST("/books", bookController.CreateBook)
		bookRoutes.GET("/books/:id", bookController.GetBook)
		bookRoutes.PUT("/books/:id", bookController.UpdateBook)
		bookRoutes.DELETE("/books/:id", bookController.DeleteBook)
	}
}
