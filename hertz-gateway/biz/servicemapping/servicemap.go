package servicemapping

import (
	"os"
	"log"
	"strings"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

// Create a service map with the service name retrieved from IDL file 
type ServiceMap struct {
	Map map[string]genericclient.Client
}

func (sm *ServiceMap) Combine(idlPath string, opts ...client.Option) {
	// Read the idl path directory and return all directory entries in it, sorted by file name
	c, err := os.ReadDir(idlPath)
	if err != nil {
		log.Fatalf("new thrift file provider failed: %v", err)
	}

	for _, entry := range c {
		if entry.IsDir() {
			sm.Combine(idlPath + entry.Name() + "/", opts...)
		} else {
			sm.addService(entry.Name(), idlPath, opts...)
		}
	}
}

func (sm *ServiceMap) addService(idlFileName string, idlPath string, opts ...client.Option) {
	
	svcName := strings.ReplaceAll(idlFileName, ".thrift", "")

	p, err := generic.NewThriftFileProvider(idlFileName, idlPath)
	if err != nil {
		panic(err)
	}

	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	cli, err := genericclient.NewClient(
		svcName,
		g,
		opts...,
	)
	if err != nil {
		panic(err)
	}

	if sm.Map == nil {
		sm.Map = make(map[string]genericclient.Client)
	}
	sm.Map[svcName] = cli
}