package handlers

import (
	"context"
	"net/http"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego-api-gateway/model/model.go"
	"errors"
)

func getBooks(ctx context.Context, c *app.RequestContext) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, utils.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}