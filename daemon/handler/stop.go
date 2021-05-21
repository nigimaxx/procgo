package handler

import (
	"context"

	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProcgoServer) Stop(_ context.Context, definitions *proto.Services) (*emptypb.Empty, error) {
	for _, svcDef := range definitions.Services {
		for _, service := range s.Services {
			if svcDef.Name == service.Name {
				close(service.StopChan)
			}
		}
	}

	return &emptypb.Empty{}, nil
}
