package xerr

import (
	"fmt"
)

func NewCodeErr(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}
func NewMsgErr(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_INTERNAL_ERROR, errMsg: errMsg}
}
func NewCodeMsgErr(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

// CodeError 常用通用固定错误
type CodeError struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}
