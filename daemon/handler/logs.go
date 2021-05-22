package handler

import (
	"github.com/nigimaxx/procgo/daemon/pkg"
	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func listenToService(svc *pkg.Service, stream proto.Procgo_LogsServer, errChan chan error) {
	logChan := make(chan []byte)
	svc.AddListener(logChan)

	for {
		line, ok := <-logChan
		if !ok {
			errChan <- nil
			break
		}

		if err := stream.Send(&wrapperspb.BytesValue{Value: line}); err != nil {
			errChan <- err
		}
	}
}

func (s *ProcgoServer) Logs(definitions *proto.AllOrServices, stream proto.Procgo_LogsServer) error {
	// is used as done channel as well if error is nil
	errChan := make(chan error)

	services := []*pkg.Service{}
	serviceCount := len(services)

	for _, svc := range s.Services {
		if definitions.All || pkg.InServiceDefList(definitions.Services, svc.Name) {
			services = append(services, svc)
			go listenToService(svc, stream, errChan)
		}
	}

	go func() {
		for {
			svc := <-s.NewServiceChan
			if (definitions.All || pkg.InServiceDefList(definitions.Services, svc.Name)) && !pkg.InServiceList(services, svc.Name) {
				serviceCount++
				go listenToService(svc, stream, errChan)
			}
		}
	}()

	for {
		err := <-errChan
		if err != nil {
			return err
		}

		serviceCount--
		if serviceCount == 0 {
			return nil
		}
	}
}
