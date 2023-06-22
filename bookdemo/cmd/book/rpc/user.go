package rpc

import (
	"context"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/user"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/user/userservice"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/constants"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/errno"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"

	// "go.opentelemetry.io/otel/metric/global"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{constants.ETCDAddress})
	if err != nil {
		panic(err)
	}
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.BookServiceName),
		provider.WithExportEndpoint(constants.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer func(ctx context.Context, p provider.OtelProvider) {
		_ = p.Shutdown(ctx)
	}(context.Background(), p)
	c, err := userservice.NewClient(
		constants.UserServiceName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.BookServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// MGetUser multiple get list of user info
func MGetUser(ctx context.Context, req *user.MGetUserRequest) (map[int64]*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := make(map[int64]*user.User)
	for _, u := range resp.Users {
		res[u.UserId] = u
	}
	return res, nil
}