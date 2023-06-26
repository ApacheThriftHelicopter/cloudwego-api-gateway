package main

import (
	management "example_hertz_kitex/kitex/kitex_gen/student/management/studentmanagement"
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

	svr := management.NewServer(new(StudentManagementImpl), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "student"}), server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
