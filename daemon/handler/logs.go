package handler

import (
	"bufio"
	"io"

	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *ProcgoServer) Logs(definitions *proto.AllOrServices, stream proto.Procgo_LogsServer) error {
	doneChan := make(chan struct{})
	errChan := make(chan error)
	serviceCount := len(s.Services)

	for _, svc := range s.Services {
		go func(svc pkg.Service) {
			// multireader
			reader := bufio.NewReader(svc.LogsReader)

			for {
				// TODO: isPrefix ???
				line, _, err := reader.ReadLine()
				if err == io.EOF {
					doneChan <- struct{}{}
					break
				}

				if err != nil {
					errChan <- err
					break
				}

				stream.Send(&wrapperspb.BytesValue{Value: append(line, []byte("\n")...)})
			}
		}(svc)
	}

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
