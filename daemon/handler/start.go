package handler

import (
	"context"

	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProcgoServer) Start(_ context.Context, definitions *proto.Services) (*emptypb.Empty, error) {
	for _, svcDef := range definitions.Services {
		svc := pkg.NewServiceFromDef(svcDef)

		if !pkg.InServiceList(s.Services, svc.Name) {
			go func(svc pkg.Service) {
				s.Services = append(s.Services, svc)

				if err := svc.Start(s.KillChan); err != nil {
					s.ErrChan <- err
					return
				}

				// remove
				for i, service := range s.Services {
					if service.Name == svc.Name {
						s.Services = append(s.Services[:i], s.Services[i+1:]...)
						s.DoneChan <- struct{}{}
						break
					}
				}

			}(svc)
		}
	}

	return &emptypb.Empty{}, nil
}
