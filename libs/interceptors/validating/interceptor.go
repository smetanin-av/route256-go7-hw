package validating

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/libs/logging"
)

type Validator interface {
	Validate() error
}

func Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if val, ok := req.(Validator); ok {
		if err := val.Validate(); err != nil {
			logging.WithCtx(ctx).Errorf("method: %q, validate: '%v'\n", info.FullMethod, err)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
	}

	return handler(ctx, req)
}
