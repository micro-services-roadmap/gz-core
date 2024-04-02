package result

import (
	"fmt"
	"github.com/micro-services-roadmap/gz-template/common/xerr"
	"github.com/pkg/errors"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// HttpResp http返回
func HttpResp(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		httpx.WriteJson(w, http.StatusOK, Success(resp))
		return
	}

	errCode := xerr.SERVER_INTERNAL_ERROR
	errMsg := xerr.MapErrMsg(xerr.SERVER_INTERNAL_ERROR)
	causeErr := errors.Cause(err)                // err类型
	if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
		errCode = e.GetErrCode()
		errMsg = e.GetErrMsg()
	} else {
		if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
			grpcCode := uint32(gstatus.Code())
			if xerr.IsCodeErr(grpcCode) { // 区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
				errCode = grpcCode
				errMsg = gstatus.Message()
			}
		}
	}

	logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
	httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
}

// ParamError 参数错误返回
func ParamError(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.REUQEST_PARAM_ERROR, errMsg))
}
