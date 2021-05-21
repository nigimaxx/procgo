package handler

import (
	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func listenToService(svc *pkg.Service, stream proto.Procgo_LogsServer, doneChan chan struct{}) {
	logChan := make(chan []byte)
	svc.AddListener(logChan)

	for {
		line, ok := <-logChan
		if !ok {
			doneChan <- struct{}{}
			break
		}

		stream.Send(&wrapperspb.BytesValue{Value: line})
	}
}

func (s *ProcgoServer) Logs(definitions *proto.AllOrServices, stream proto.Procgo_LogsServer) error {
	doneChan := make(chan struct{})
	errChan := make(chan error)

	services := []*pkg.Service{}
	serviceCount := len(services)

	for _, svc := range s.Services {
		if definitions.All || pkg.InServiceDefList(definitions.Services, svc.Name) {
			services = append(services, svc)
			go listenToService(svc, stream, doneChan)
		}
	}

	go func() {
		for {
			svc := <-s.NewServiceChan
			if (definitions.All || pkg.InServiceDefList(definitions.Services, svc.Name)) && !pkg.InServiceList(services, svc.Name) {
				serviceCount++
				go listenToService(svc, stream, doneChan)
			}
		}
	}()

	for {
		select {
		case err := <-errChan:
			return err
		case <-doneChan:
			serviceCount--
			if serviceCount == 0 {
				return nil
			}
		}
	}
}
