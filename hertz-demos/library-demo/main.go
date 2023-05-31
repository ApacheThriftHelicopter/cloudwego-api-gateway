package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/hertz-demos/library-demo/handlers"
)

func main() {
	// Set-up router
	h := server.Default()

	h.GET("/books", handlers.getBooks)
	h.GET("/books/:id", handlers.bookById)
	h.POST("/books", handlers.createBook)
	h.PATCH("/checkout", handlers.checkoutBook)
	h.PATCH("/return", handlers.returnBook)

	h.Spin()
}