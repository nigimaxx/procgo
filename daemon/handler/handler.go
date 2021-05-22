package handler

import (
	"os"

	"github.com/nigimaxx/procgo/daemon/pkg"
	"github.com/nigimaxx/procgo/proto"
)

type ProcgoServer struct {
	proto.UnimplementedProcgoServer
	Services       []*pkg.Service
	NewServiceChan chan *pkg.Service
	// ErrChan is used as done channel as well if error is nil
	ErrChan  chan error
	KillChan chan os.Signal
}

func NewProcgoServer() ProcgoServer {
	return ProcgoServer{
		Services:       []*pkg.Service{},
		NewServiceChan: make(chan *pkg.Service, 64), // max running process without logs. TODO: !!!
		ErrChan:        make(chan error),
		KillChan:       make(chan os.Signal, 1),
	}
}

func (s *ProcgoServer) startService(svc *pkg.Service) {
	s.Services = append(s.Services, svc)
	s.NewServiceChan <- svc

	if err := svc.Start(s.KillChan); err != nil {
		s.ErrChan <- err
		return
	}

	for i, service := range s.Services {
		if service.Name == svc.Name {
			s.Services = append(s.Services[:i], s.Services[i+1:]...)
			s.ErrChan <- nil
			break
		}
	}
}
