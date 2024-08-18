package middleware

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/x1rh/gopkg/jwtx"
	"github.com/x1rh/zero-contrib/gwx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Auth(jwtManager *jwtx.Manager, mapper *gwx.Router) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			app := r.Header.Get("x-app")
			logx.Debugf("app:%s\n", app)
			r.Header.Set("Grpc-Metadata-app", app)
			if !mapper.IsRequireAuth(r.Method, r.RequestURI) {
				next(w, r)
				return
			}

			authorization := r.Header.Get("authorization")
			logx.Debugf("authorization:%s\n", authorization)
			claims, err := jwtManager.WithApp(app).Verify(authorization)
			if err != nil {
				err := errors.Wrap(err, "fail to parse header `authorization`")
				logx.Error(err)
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
			r.Header.Set("Grpc-Metadata-uid", fmt.Sprintf("%d", claims.User.Uid))

			next(w, r)
		}
	}
}
