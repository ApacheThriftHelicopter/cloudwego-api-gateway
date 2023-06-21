// Code generated by hertz generator.

package main

import (
	handler "github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/bookapi/hertz_handler"
	
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/api/hertz_handler"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/api/hertz_handler/bookapi"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	r.NoRoute(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 404
		bookapi.SendResponse(c, errno.ServiceErr, nil)
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) { // used for HTTP 405
		bookapi.SendResponse(c, errno.ServiceErr, nil)
	})
}
