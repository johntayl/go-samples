package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book represents a book
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type BookRequest struct {
	ID int `uri:"id" binding:"required"`
}

// books slice to seed record data
var books = []Book{
	{ID: 1, Title: "Book 1", Author: "Author 1"},
	{ID: 2, Title: "Book 2", Author: "Author 2"},
}

func main() {
	// Create a new router
	router := gin.Default()

	// Bind the resource to the router (books)
	router.GET("/books", GetBooks)
	router.GET("/books/:id", GetBook)
	router.POST("/books", CreateBook)
	router.PUT("/books/:id", UpdateBook)
	router.DELETE("/books/:id", DeleteBook)

	// Run the server
	router.Run("localhost:8080")
}

// GetBooks responds with the list of all books as JSON.
func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// GetBook responds with the book of the given ID.
func GetBook(c *gin.Context) {
	var bookRequest BookRequest

	if err := c.BindUri(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, a := range books {
		if a.ID == bookRequest.ID {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// CreateBook adds a book from JSON received in the request body.
func CreateBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	newBook.ID = len(books) + 1

	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// UpdateBook updates a book from JSON received in the request body.
func UpdateBook(c *gin.Context) {
	var bookRequest BookRequest

	if err := c.BindUri(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range books {
		if a.ID == bookRequest.ID {
			var updateBook Book
			if err := c.BindJSON(&updateBook); err != nil {
				return
			}
			updateBook.ID = bookRequest.ID
			books[i] = updateBook
			c.JSON(http.StatusOK, updateBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// DeleteBook deletes a book by its ID.
func DeleteBook(c *gin.Context) {
	var bookRequest BookRequest

	if err := c.BindUri(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, a := range books {
		if a.ID == bookRequest.ID {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
