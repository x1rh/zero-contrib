package errx

import (
	"net/http"
	"zero-contrib/errx/types"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type codeMsgData struct {
	Code int    `json:"code" xml:"code"`
	Msg  string `json:"msg" xml:"msg"`
	Data any    `json:"data,omitempty" xml:"data,omitempty"`
}

func GrpcCodeToHttpCode(code codes.Code) int {
	switch code {
	case codes.InvalidArgument, codes.Aborted:
		return 400
	case codes.Unauthenticated:
		return 401
	case codes.PermissionDenied:
		return 403
	case codes.NotFound:
		return 404
	default:
		return 500
	}
}

func NewErrorHandler() func(err error) (int, any) {
	return func(err error) (int, any) {
		// must using package errx to wrap all errors
		switch e := err.(type) {
		case error:
			s, ok := status.FromError(e)
			if ok {
				// todo: is always a grpc error ?
				for _, d := range s.Details() {
					switch info := d.(type) {
					case *types.ErrorMessage:
						logx.Error(info.Message)
					default:
						logx.Errorf("Unexpected type: %s", info)
					}
				}

				code := GrpcCodeToHttpCode(s.Code())
				return code, codeMsgData{
					Code: code,
					Msg:  s.Message(),
				}
			} else {
				logx.Error(e.Error())
				return 500, codeMsgData{
					Code: 500,
					Msg:  e.Error(),
				}
			}
		default:
			logx.Error(err)
			return http.StatusInternalServerError, nil
		}
	}
}
