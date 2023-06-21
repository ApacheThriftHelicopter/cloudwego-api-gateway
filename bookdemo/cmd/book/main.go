package main

import (
	"context"
	"net"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/dal"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/rpc"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book/bookservice"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/constants"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	rpc.Init()
	dal.Init()
	// klog init
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(constants.TCP, constants.BookServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constants.BookServiceName),
		provider.WithExportEndpoint(constants.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer func(ctx context.Context, p provider.OtelProvider) {
		_ = p.Shutdown(ctx)
	}(context.Background(), p)
	svr := bookservice.NewServer(new(BookServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.BookServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}