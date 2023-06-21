package service

import (
	"context"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/dal/db"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/pack"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/rpc"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/user"
)

type QueryBookService struct {
	ctx context.Context
}

// NewQueryBookService new QueryBookService
func NewQueryBookService(ctx context.Context) *QueryBookService {
	return &QueryBookService{ctx: ctx}
}

// QueryBookService query list of book info
func (s *QueryBookService) QueryBookService(req *book.QueryBookRequest) ([]*book.Book, int64, error) {
	bookModels, total, err := db.QueryBook(s.ctx, req.UserId, req.SearchKey, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, 0, err
	}
	userMap, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserIds: []int64{req.UserId}})
	if err != nil {
		return nil, 0, err
	}
	books := pack.Books(bookModels)
	for i := 0; i < len(books); i++ {
		if u, ok := userMap[books[i].UserId]; ok {
			books[i].Username = u.Username
			books[i].UserAvatar = u.Avatar
		}
	}
	return books, total, nil
}