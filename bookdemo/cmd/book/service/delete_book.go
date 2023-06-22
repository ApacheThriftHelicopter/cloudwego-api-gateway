package service

import (
	"context"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/dal/db"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
)

type DelBookService struct {
	ctx context.Context
}

// NewDelBookService new DelBookService
func NewDelBookService(ctx context.Context) *DelBookService {
	return &DelBookService{
		ctx: ctx,
	}
}

// DelBook delete book info
func (s *DelBookService) DelBook(req *book.DeleteBookRequest) error {
	return db.DeleteBook(s.ctx, req.BookId, req.UserId)
}