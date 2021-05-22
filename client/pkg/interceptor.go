package pkg

import (
	"context"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

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
