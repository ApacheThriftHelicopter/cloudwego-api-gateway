package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// Set-up router
	h := server.Default()

	h.GET("/books", getBooks)
	h.GET("/books/:id", bookById)
	h.POST("/books", createBook)
	h.PATCH("/checkout", checkoutBook)
	h.PATCH("/return", returnBook)

	h.Spin()
}