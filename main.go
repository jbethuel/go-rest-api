package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var books = []book{
	{
		Id:        "1",
		Name:      "A Brief History of Time",
		Completed: false,
	},
	{
		Id:        "2",
		Name:      "The Hitchhiker's Guide to the Galaxy",
		Completed: false,
	},
	{
		Id:        "3",
		Name:      "Berserk",
		Completed: false,
	},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:9090")
}
