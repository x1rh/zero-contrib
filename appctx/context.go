package zerox

import (
	"context"

	"github.com/x1rh/zero-contrib/interceptorx"
)

type AppContext struct {
	Uid int64
}

func GetContext(ctx context.Context) *AppContext{
	uid := ctx.Value(interceptorx.UID).(int64)

	return &AppContext{
		Uid: uid,
	}
}
