package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var books = []Book{
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

func getBookById(id string) (book Book, err error) {
	for i, b := range books {
		if b.Id == id {
			return books[i], nil
		}
	}

	return Book{Id: "", Name: "", Completed: false}, errors.New("Document not found")
}

func getBook(context *gin.Context) {
	id := context.Param("id")
	book, err := getBookById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	context.IndentedJSON(http.StatusOK, book)
}

func addBook(context *gin.Context) {
	var newBook Book

	err := context.BindJSON(&newBook)
	if err != nil {
		return
	}

	books = append(books, newBook)

	context.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()

	router.GET("/books/:id", getBook)
	router.GET("/books", getBooks)
	router.POST("/books", addBook)

	router.Run("localhost:9090")
}
