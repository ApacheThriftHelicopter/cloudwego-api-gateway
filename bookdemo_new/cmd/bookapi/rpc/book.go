package rpc

import (
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/constants"
	
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/cloudwego/kitex/pkg/generic"
    "github.com/cloudwego/kitex/client/genericclient"
)

var genericBookClient genericclient.Client

func initBook() {
	r, err := etcd.NewEtcdResolver([]string{constants.ETCDAddress})
	if err != nil {
		panic(err)
	}

	// Parse IDL with Local Files
    // YOUR_IDL_PATH thrift file path, eg:./idl/example.thrift
    t, err := generic.NewThriftFileProvider("../../../idl/book.thrift")
    if err != nil {
        panic(err)
    }
    g, err := generic.JSONThriftGeneric(t)
    if err != nil {
        panic(err)
    }
	c, err := genericclient.NewClient(
		constants.BookServiceName,
		g,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	genericBookClient = c
}
