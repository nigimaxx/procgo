package handler

import (
	"context"

	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProcgoServer) Restart(_ context.Context, definitions *proto.Services) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
