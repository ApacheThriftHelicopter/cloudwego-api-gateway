package handlers

import (
	"context"
	"net/http"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego-api-gateway/model/model.go"
)

func createBook(ctx context.Context, c *app.RequestContext) {
	var newBook book 

	if err := c.Bind(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}