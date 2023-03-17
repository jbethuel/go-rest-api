package books

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

func remove(slice []Book, s int) []Book {
	return append(slice[:s], slice[s+1:]...)
}

func deleteBookById(id string) (bool, error) {
	var index int = -1
	for i, b := range books {
		if b.Id == id {
			index = i
		}
	}
	if index == -1 {
		return false, errors.New("Document not found")
	}

	books = remove(books, index)

	return true, nil
}

func getBookById(id string) (Book, error) {
	for i, b := range books {
		if b.Id == id {
			return books[i], nil
		}
	}

	return Book{Id: "", Name: "", Completed: false}, errors.New("Document not found")
}

func checkBook(id string) bool {
	for _, b := range books {
		if b.Id == id {
			return true
		}
	}

	return false
}

func GetBook(context *gin.Context) {
	id := context.Param("id")
	book, err := getBookById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	context.IndentedJSON(http.StatusOK, book)
}

func DeleteBook(context *gin.Context) {
	id := context.Param("id")

	_, err := deleteBookById(id)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, "success")
}

func GetBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func AddBook(context *gin.Context) {
	var newBook Book

	err := context.BindJSON(&newBook)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	doesBookExists := checkBook(newBook.Id)
	if doesBookExists == true {
		context.IndentedJSON(http.StatusBadRequest, "Book already added")
		return
	}

	books = append(books, newBook)

	context.IndentedJSON(http.StatusCreated, newBook)
}
