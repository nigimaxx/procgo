package pkg

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func WithServerUnaryInterceptor(procfile string) grpc.ServerOption {
	return grpc.UnaryInterceptor(func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
		}

		procfileHeader, ok := md["procfile"]
		if !ok || len(procfileHeader) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "Procfile not supplied")
		}

		if procfile != procfileHeader[0] {
			log.Println(procfile, procfileHeader)
			return nil, status.Errorf(codes.InvalidArgument, "Wrong Procfile")
		}

		h, err := handler(ctx, req)

		return h, err
	})
}
