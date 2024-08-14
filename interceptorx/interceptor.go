package interceptorx

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type contextKey int

const (
	IP contextKey = iota
	APP
	UID
)

func MetadataInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		var err error

		var ip string
		var uid int64
		var app string

		ipList := md.Get("gateway-ip")
		if len(ipList) >= 1 {
			ip = ipList[0]
		}

		appList := md.Get("gateway-app")
		if len(appList) >= 1 {
			app = appList[0]
		}

		uidList := md.Get("gateway-uid")
		if len(uidList) >= 1 {
			uid, err = strconv.ParseInt(uidList[0], 10, 64)
			if err != nil {
				logx.Error(err)
			}
		}

		// ctx = context.WithValue(ctx, "ip", ip)
		// ctx = context.WithValue(ctx, "app", app)
		// ctx = context.WithValue(ctx, "uid", uid)
		ctx = context.WithValue(ctx, IP, ip)
		ctx = context.WithValue(ctx, APP, app)
		ctx = context.WithValue(ctx, UID, uid)
	}

	return handler(ctx, req)
}
