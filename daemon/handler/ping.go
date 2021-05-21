package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProcgoServer) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
