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

var blankBook = Book{Id: "", Name: "", Completed: false}

func remove(slice []Book, s int) []Book {
	return append(slice[:s], slice[s+1:]...)
}

func checkBook(id string) bool {
	for _, b := range books {
		if b.Id == id {
			return true
		}
	}

	return false
}

func addBookToBooks(book Book) error {
	doesBookExists := checkBook(book.Id)
	if doesBookExists == true {
		return errors.New("Book already added")
	}

	books = append(books, book)

	return nil
}

func deleteBookById(id string) error {
	doesBookExists := checkBook(id)
	if doesBookExists == false {
		return errors.New("Document not found")
	}

	var index int = -1
	for i, b := range books {
		if b.Id == id {
			index = i
		}
	}
	if index == -1 {
		return errors.New("Unexpected error occured.")
	}

	books = remove(books, index)

	return nil
}

func getBookById(id string) (Book, error) {
	doesBookExists := checkBook(id)
	if doesBookExists == false {
		return blankBook, errors.New("Document not found")
	}

	for i, b := range books {
		if b.Id == id {
			return books[i], nil
		}
	}

	return blankBook, errors.New("Unexpected error occured.")
}

func updateBookById(id string, b Book) error {
	doesBookExists := checkBook(id)
	if doesBookExists == false {
		return errors.New("Document not found")
	}

	for i, eachBook := range books {
		if eachBook.Id == id {
			books[i] = b
			return nil
		}
	}

	return errors.New("Unexpected error occured.")
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

	err := deleteBookById(id)
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

	err1 := context.BindJSON(&newBook)
	if err1 != nil {
		context.IndentedJSON(http.StatusBadRequest, err1.Error())
		return
	}

	err2 := addBookToBooks(newBook)
	if err2 != nil {
		context.IndentedJSON(http.StatusBadRequest, err2.Error())
		return
	}

	context.IndentedJSON(http.StatusCreated, newBook)
}

func PatchBook(context *gin.Context) {
	id := context.Param("id")

	var newBook Book

	err1 := context.BindJSON(&newBook)
	if err1 != nil {
		context.IndentedJSON(http.StatusBadRequest, err1.Error())
		return
	}

	err2 := updateBookById(id, newBook)
	if err2 != nil {
		context.IndentedJSON(http.StatusBadRequest, err2.Error())
		return
	}

	context.IndentedJSON(http.StatusOK, newBook)
}
