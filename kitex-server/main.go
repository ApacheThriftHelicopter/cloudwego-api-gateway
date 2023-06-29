package main

import (
	book "github.com/ApacheThriftHelicopter/cloudwego-api-gateway/kitex-server/kitex_gen/book/booksvc"
	"log"
	
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
    if err != nil {
        panic(err)
    }

	svr := book.NewServer(new(BookSvcImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "book"}), server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
