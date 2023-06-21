package service

import (
	"context"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/dal/db"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
)

type UpdateBookService struct {
	ctx context.Context
}

// NewUpdateBookService new UpdateBookService
func NewUpdateBookService(ctx context.Context) *UpdateBookService {
	return &UpdateBookService{ctx: ctx}
}

// UpdateBook update book info
func (s *UpdateBookService) UpdateBook(req *book.UpdateBookRequest) error {
	return db.UpdateBook(s.ctx, req.BookId, req.UserId, req.Title, req.Content)
}