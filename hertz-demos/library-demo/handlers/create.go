package handlers

import (
	"context"
	"net/http"
	
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/hertz-demos/library-demo/model"
)

func createBook(ctx context.Context, c *app.RequestContext) {
	var newBook model.book 

	if err := c.Bind(&newBook); err != nil {
		return
	}

	model.books = append(model.books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}