package rpc

import (
	"context"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book/bookservice"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/constants"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/errno"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var bookClient bookservice.Client

func initBook() {
	r, err := etcd.NewEtcdResolver([]string{constants.ETCDAddress})
	if err != nil {
		panic(err)
	}
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.ApiServiceName),
		provider.WithExportEndpoint(constants.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer func(ctx context.Context, p provider.OtelProvider) {
		_ = p.Shutdown(ctx)
	}(context.Background(), p)
	c, err := bookservice.NewClient(
		constants.BookServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	bookClient = c
}

// CreateBook create book info
func CreateBook(ctx context.Context, req *book.CreateBookRequest) error {
	resp, err := bookClient.CreateBook(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// QueryBooks query list of book info
func QueryBooks(ctx context.Context, req *book.QueryBookRequest) ([]*book.Book, int64, error) {
	resp, err := bookClient.QueryBook(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Books, resp.Total, nil
}

// UpdateBook update book info
func UpdateBook(ctx context.Context, req *book.UpdateBookRequest) error {
	resp, err := bookClient.UpdateBook(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// DeleteBook delete book info
func DeleteBook(ctx context.Context, req *book.DeleteBookRequest) error {
	resp, err := bookClient.DeleteBook(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}