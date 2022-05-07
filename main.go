package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)
type book struct {
	ID       string `json: "id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getbook(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, books)

}
func AddBook(c *gin.Context) {
	var newbook book
	if err := c.BindJSON(&newbook); err != nil {
		return
	}
	books = append(books, newbook)

}

func getbooks(c *gin.Context) {

	id := c.Param("id")
	newbookpointer, err := gogetbyid(id)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, newbookpointer)

}

func gogetbyid(id string) (*book, error) {

	for i, b := range books {
		if b.ID == id {
			return &books[i], nil

		}

	}
	return nil, errors.New("deepak ")

}
func main() {

	route := gin.Default()
	route.GET("/BOOK", getbook)
	route.POST("/AddBook", AddBook)
	route.GET("/Book/:id", getbooks)
	route.Run("localhost:8080")

}
