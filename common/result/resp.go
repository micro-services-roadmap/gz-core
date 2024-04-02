package result

import "github.com/micro-services-roadmap/gz-template/common/xerr"

type NullJson struct{}

type Response[T any] struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func Success[T any](data T) *Response[T] {
	return &Response[T]{xerr.OK, xerr.MapErrMsg(xerr.OK), data}
}

func ErrorCode(errCode uint32) *Response[any] {

	return &Response[any]{Code: errCode, Msg: xerr.MapErrMsg(errCode)}
}

func Error(errCode uint32, errMsg string) *Response[any] {

	return &Response[any]{Code: errCode, Msg: errMsg}
}
