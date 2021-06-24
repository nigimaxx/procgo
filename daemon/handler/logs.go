package handler

import (
	"sync"
	"time"

	"github.com/nigimaxx/procgo/daemon/pkg"
	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type logListener struct {
	mu       sync.Mutex
	services map[string]time.Time
	// is used as done channel as well if error is nil
	errChan chan error
}

func (l *logListener) listenToService(svc *pkg.Service, stream proto.Procgo_LogsServer) {
	logChan := make(chan []byte)
	svc.AddListener(logChan)

	for {
		line, ok := <-logChan
		if !ok {
			l.removeService(svc.Name)
			l.errChan <- nil
			break
		}

		if err := stream.Send(&wrapperspb.BytesValue{Value: line}); err != nil {
			l.errChan <- err
		}
	}
}

func (l *logListener) addService(name string) {
	l.mu.Lock()
	l.services[name] = time.Now()
	l.mu.Unlock()
}

func (l *logListener) removeService(name string) {
	l.mu.Lock()
	t, ok := l.services[name]
	if ok && time.Since(t) > 1*time.Second {
		delete(l.services, name)
	}
	l.mu.Unlock()
}

func (s *ProcgoServer) Logs(definitions *proto.AllOrServices, stream proto.Procgo_LogsServer) error {
	listener := logListener{services: make(map[string]time.Time), errChan: make(chan error)}

	for _, svc := range s.Services {
		if definitions.All || pkg.InServiceDefList(definitions.Services, svc.Name) {

			listener.addService(svc.Name)
			go listener.listenToService(svc, stream)
		}
	}

	go func() {
		for {
			svc := <-s.NewServiceChan

			listener.mu.Lock()
			t, ok := listener.services[svc.Name]
			listener.mu.Unlock()

			if (definitions.All || pkg.InServiceDefList(definitions.Services, svc.Name)) && (!ok || time.Since(t) > 1*time.Second) {
				listener.addService(svc.Name)
				go listener.listenToService(svc, stream)
			}
		}
	}()

	for {
		err := <-listener.errChan
		if err != nil {
			return err
		}

		listener.mu.Lock()
		length := len(listener.services)
		listener.mu.Unlock()

		if length == 0 {
			return nil
		}
	}
}
