package service

import (
	"context"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/dal/db"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/pack"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/rpc"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/user"
)

type MGetBookService struct {
	ctx context.Context
}

// NewMGetBookService new MGetBookService
func NewMGetBookService(ctx context.Context) *MGetBookService {
	return &MGetBookService{ctx: ctx}
}

// MGetBook multiple get list of book info
func (s *MGetBookService) MGetBook(req *book.MGetBookRequest) ([]*book.Book, error) {
	bookModels, err := db.MGetBooks(s.ctx, req.BookIds)
	if err != nil {
		return nil, err
	}
	uIds := pack.UserIds(bookModels)
	userMap, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserIds: uIds})
	if err != nil {
		return nil, err
	}
	books := pack.Books(bookModels)
	for i := 0; i < len(books); i++ {
		if u, ok := userMap[books[i].UserId]; ok {
			books[i].Username = u.Username
			books[i].UserAvatar = u.Avatar
		}
	}
	return books, nil
}