package pkg

import (
	"context"
	"log"
	"path/filepath"

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

func WithClientUnaryInterceptor(procfile string) grpc.DialOption {
	return grpc.WithUnaryInterceptor(func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		absProcfile, err := filepath.Abs(procfile)
		if err != nil {
			return err
		}

		md := metadata.New(map[string]string{"procfile": absProcfile})
		withMetatdata := metadata.NewOutgoingContext(ctx, md)

		return invoker(withMetatdata, method, req, reply, cc, opts...)
	})
}
