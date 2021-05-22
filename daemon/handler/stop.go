package handler

import (
	"context"

	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProcgoServer) Stop(_ context.Context, definitions *proto.Services) (*emptypb.Empty, error) {
	for _, svcDef := range definitions.Services {
		for i, service := range s.Services {
			if svcDef.Name == service.Name {
				s.Services = append(s.Services[:i], s.Services[i+1:]...)
				close(service.StopChan)
			}
		}
	}

	s.ErrChan <- nil

	return &emptypb.Empty{}, nil
}
