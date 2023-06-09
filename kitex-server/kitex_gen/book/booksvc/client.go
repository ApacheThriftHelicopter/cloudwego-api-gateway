// Code generated by Kitex v0.5.2. DO NOT EDIT.

package booksvc

import (
	"context"
	book "github.com/ApacheThriftHelicopter/cloudwego-api-gateway/kitex-server/kitex_gen/book"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	QueryBook(ctx context.Context, req *book.QueryBookReq, callOptions ...callopt.Option) (r *book.QueryBookResp, err error)
	InsertBook(ctx context.Context, req *book.InsertBookReq, callOptions ...callopt.Option) (r *book.InsertBookResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kBookSvcClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kBookSvcClient struct {
	*kClient
}

func (p *kBookSvcClient) QueryBook(ctx context.Context, req *book.QueryBookReq, callOptions ...callopt.Option) (r *book.QueryBookResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryBook(ctx, req)
}

func (p *kBookSvcClient) InsertBook(ctx context.Context, req *book.InsertBookReq, callOptions ...callopt.Option) (r *book.InsertBookResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InsertBook(ctx, req)
}
