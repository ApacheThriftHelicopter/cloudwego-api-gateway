// Code generated by hertz generator. DO NOT EDIT.

package main

import (
	router "example_hertz_kitex/hertz/biz/router"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
