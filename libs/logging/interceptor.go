package logging

import (
	"context"

	"google.golang.org/grpc"
)

func Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res any, err error) {
	WithCtx(ctx).Infof("method: %q, request: '%v'\n", info.FullMethod, req)
	res, err = handler(ctx, req)
	if err == nil {
		WithCtx(ctx).Infof("method: %q, response: '%v'\n", info.FullMethod, res)
	} else {
		WithCtx(ctx).Errorf("method: %q, error: '%v'\n", info.FullMethod, err)
	}
	return res, err
}
