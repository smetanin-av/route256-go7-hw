package panic_recover

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res any, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			log.Printf("FATAL method: %q, panic: %v\n", info.FullMethod, msg)
			err = status.Errorf(codes.Internal, "panic: %v", msg)
		}
	}()

	return handler(ctx, req)
}
