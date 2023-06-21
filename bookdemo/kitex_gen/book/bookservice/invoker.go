// Code generated by Kitex v0.6.1. DO NOT EDIT.

package bookservice

import (
	book "github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler book.BookService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
