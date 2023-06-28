package main

import (
	"context"
	book "github.com/ApacheThriftHelicopter/cloudwego-api-gateway/kitex-server/kitex_gen/book"
)

// BookSvcImpl implements the last service interface defined in the IDL.
type BookSvcImpl struct{}

// For QueryBookResp, each book contains these info
type BookInfo struct {
	ID string
	Title string
	Author string
}

// The demo server stores 5 book info
var BookData = make(map[string]BookInfo, 5)

// QueryBook implements the BookSvcImpl interface.
func (s *BookSvcImpl) QueryBook(ctx context.Context, req *book.QueryBookReq) (resp *book.QueryBookResp, err error) {
	num := string(req.Num)
	b, exist := BookData[num]

	if !exist {
		return &book.QueryBookResp{
			Exist: false,
		}, nil
	}

	resp = &book.QueryBookResp{
		Exist: true,
		ID: b.ID,
		Title: b.Title,
		Author: b.Author,
	}

	return resp, nil
}

// InsertBook implements the BookSvcImpl interface.
func (s *BookSvcImpl) InsertBook(ctx context.Context, req *book.InsertBookReq) (resp *book.InsertBookResp, err error) {
	_, exist := BookData[req.ID]
	
	// Unsuccessful response
	if exist {
		return &book.InsertBookResp{
			Ok: false,
			Msg: "There is a book with that ID already.",
		}, nil
	}

	// Insert if does not exist
	BookData[req.ID] = BookInfo{
		ID: req.ID,
		Title: req.Title,
		Author: req.Author,
	}

	// Successful response
	resp = &book.InsertBookResp{
		Ok: true,
		Msg: "Book successfully inserted.",
	}

	return resp, nil
}
