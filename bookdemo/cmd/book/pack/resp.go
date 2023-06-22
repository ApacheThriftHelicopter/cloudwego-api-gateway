package pack

import (
	"errors"
	"time"

	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/kitex_gen/book"
	"github.com/ApacheThriftHelicopter/cloudwego-api-gateway/bookdemo/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *book.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *book.BaseResp {
	return &book.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}