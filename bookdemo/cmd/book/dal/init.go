package dal

import "github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/cmd/book/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}