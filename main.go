package main

import (
	books "go-rest-api/books"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books/:id", books.GetBook)
	router.DELETE("/books/:id", books.DeleteBook)
	router.GET("/books", books.GetBooks)
	router.POST("/books", books.AddBook)

	router.Run("localhost:9090")
}
