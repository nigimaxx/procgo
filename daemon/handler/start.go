package handler

import (
	"context"

	"github.com/nigimaxx/procgo/daemon/pkg"
	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProcgoServer) Start(_ context.Context, definitions *proto.Services) (*emptypb.Empty, error) {
	for _, svcDef := range definitions.Services {
		svc := pkg.NewServiceFromDef(svcDef)

		if !pkg.InServiceList(s.Services, svc.Name) {
			go s.startService(svc)
		}
	}

	return &emptypb.Empty{}, nil
}
