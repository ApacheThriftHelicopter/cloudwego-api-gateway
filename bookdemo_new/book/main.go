package main

import (
	"context"
	"fmt"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	
	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server/genericserver"
)

type GenericServiceImpl struct{}

func (s *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	// use jsoniter or other json parse sdk to assert request
	switch method {
	case "CreateBook":
		return 
	case "QueryBook":
		return
	default:
		return nil, fmt.Errorf("unknown method: %s", method)
	}
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
	
	// Parse IDL with Local Files
    // YOUR_IDL_PATH thrift file path,eg: ./idl/example.thrift
    t, err := generic.NewThriftFileProvider("../../idl/book.thrift")
    if err != nil {
        panic(err)
    }
    g, err := generic.JSONThriftGeneric(t)
    if err != nil {
        panic(err)
    }
	svr := genericserver.NewServer(new(GenericServiceImpl),
		g,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.BookServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}