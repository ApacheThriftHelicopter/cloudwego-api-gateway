package handlers

import (
	"context"
	"net/http"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/hertz-demos/library-demo/model"
)

func getBooks(ctx context.Context, c *app.RequestContext) {
	c.IndentedJSON(http.StatusOK, model.books)
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

func getBookById(id string) (*model.book, error) {
	for i, b := range model.books {
		if b.ID == id {
			return &model.books[i], nil
		}
	}

	return nil, errors.New("book not found")
}