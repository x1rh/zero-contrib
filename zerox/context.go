package zerox

import (
	"context"
	"zero-contrib/zerox/interceptorx"
)

type ZeroContext struct {
	Uid int64
}

func GetContext(ctx context.Context) *ZeroContext {
	uid := ctx.Value(interceptorx.UID).(int64)

	return &ZeroContext{
		Uid: uid,
	}
}
