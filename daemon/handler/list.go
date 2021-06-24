package handler

import (
	"context"

	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProcgoServer) List(_ context.Context, _ *emptypb.Empty) (*proto.Services, error) {
	services := []*proto.ServiceDefinition{}

	for _, svc := range s.Services {
		services = append(services, &proto.ServiceDefinition{
			Name:    svc.Name,
			Command: svc.Command,
		})
	}

	return &proto.Services{Services: services}, nil
}
