package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProcgoServer) KillAll(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	close(s.KillChan)
	s.ErrChan <- nil
	return &emptypb.Empty{}, nil
}
