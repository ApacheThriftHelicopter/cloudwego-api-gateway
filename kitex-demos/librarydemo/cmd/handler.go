package main

import (
	librarydemo "book/idl/kitex_gen/librarydemo"
	"context"
)

// BookServiceImpl implements the last service interface defined in the IDL.
type BookServiceImpl struct{}

// CreateBook implements the BookServiceImpl interface.
func (s *BookServiceImpl) CreateBook(ctx context.Context, req *librarydemo.CreateBookRequest) (resp *librarydemo.CreateBookResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetBook implements the BookServiceImpl interface.
func (s *BookServiceImpl) MGetBook(ctx context.Context, req *librarydemo.MGetBookRequest) (resp *librarydemo.MGetBookResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteBook implements the BookServiceImpl interface.
func (s *BookServiceImpl) DeleteBook(ctx context.Context, req *librarydemo.DeleteNoteRequest) (resp *librarydemo.DeleteBookResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryBook implements the BookServiceImpl interface.
func (s *BookServiceImpl) QueryBook(ctx context.Context, req *librarydemo.QueryBookRequest) (resp *librarydemo.QueryBookResponse, err error) {
	// TODO: Your code here...
	return
}
