package main

import (
	"context"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/pack"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/service"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/constants"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/errno"
)

// BookServiceImpl implements the last service interface defined in the IDL.
type BookServiceImpl struct{}

// CreateBook implements the BookServiceImpl interface.
func (s *BookServiceImpl) CreateBook(ctx context.Context, req *book.CreateBookRequest) (resp *book.CreateBookResponse, err error) {
	resp = new(book.CreateBookResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateBookService(ctx).CreateBook(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteBook implements the BookServiceImpl interface.
func (s *BookServiceImpl) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (resp *book.DeleteBookResponse, err error) {
	resp = new(book.DeleteBookResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewDelBookService(ctx).DelBook(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// UpdateBook implements the BookServiceImpl interface.
func (s *BookServiceImpl) UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (resp *book.UpdateBookResponse, err error) {
	resp = new(book.UpdateBookResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateBookService(ctx).UpdateBook(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// QueryBook implements the BookServiceImpl interface.
func (s *BookServiceImpl) QueryBook(ctx context.Context, req *book.QueryBookRequest) (resp *book.QueryBookResponse, err error) {
	resp = new(book.QueryBookResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	if req.Limit == 0 {
		req.Limit = consts.DefaultLimit
	}

	books, total, err := service.NewQueryBookService(ctx).QueryBookService(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Books = books
	resp.Total = total
	return resp, nil
}

// MGetBook implements the BookServiceImpl interface.
func (s *BookServiceImpl) MGetBook(ctx context.Context, req *book.MGetBookRequest) (resp *book.MGetBookResponse, err error) {
	resp = new(book.MGetBookResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	books, err := service.NewMGetBookService(ctx).MGetBook(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Books = books
	return resp, nil
}