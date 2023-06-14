package main

import (
	librarydemo "book/idl/kitex_gen/librarydemo/bookservice"
	"log"
)

func main() {
	svr := librarydemo.NewServer(new(BookServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
