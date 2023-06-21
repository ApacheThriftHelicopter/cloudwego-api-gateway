package service

import (
	"context"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/dal/db"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
)

type CreateBookService struct {
	ctx context.Context
}

// NewCreateBookService new CreateBookService
func NewCreateBookService(ctx context.Context) *CreateBookService {
	return &CreateBookService{ctx: ctx}
}

// CreateBook create book info
func (s *CreateBookService) CreateBook(req *book.CreateBookRequest) error {
	bookModel := &db.Book{
		UserID:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	}
	return db.CreateBook(s.ctx, []*db.Book{bookModel})
}