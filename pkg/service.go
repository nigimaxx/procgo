package pkg

import (
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/fatih/color"
	"github.com/nigimaxx/procgo/proto"
)

type Service struct {
	Name    string
	Command string

	StopChan chan os.Signal
	// currently if no one is listening the writer is blocking
	// possible problem with broadcasting logs to multiple instances of the client
	LogsReader *io.PipeReader
	LogsWriter *io.PipeWriter
}

var shell string

func init() {
	shell = os.Getenv("SHELL")
	if shell == "" {
		panic("$SHELL not defined")
	}
}

func NewServiceFromDef(s *proto.ServiceDefinition) Service {
	r, w := io.Pipe()
	return Service{s.Name, s.Command, make(chan os.Signal), r, w}
}

func (s *Service) ToDef() *proto.ServiceDefinition {
	return &proto.ServiceDefinition{Name: s.Name, Command: s.Command}
}

func (s *Service) Start(killChan <-chan os.Signal) error {
	cmd := exec.Command(shell, "-c", s.Command)

	out := NewPrefixWriter(s.Name, s.LogsWriter)

	cmd.Stdout = out
	cmd.Stderr = out

	if err := cmd.Start(); err != nil {
		return err
	}

	cmdErrChan := make(chan error)
	doneChan := make(chan struct{})

	go func() {
		err := cmd.Wait()
		if err != nil {
			cmdErrChan <- err
		} else {
			doneChan <- struct{}{}
		}
	}()

	select {
	case <-killChan:
		out.Write([]byte(color.YellowString("Killing with signal %v", syscall.SIGINT)))
		cmd.Process.Signal(syscall.SIGINT)
		return s.LogsWriter.Close()
	case <-s.StopChan:
		out.Write([]byte(color.YellowString("Killing with signal %v", syscall.SIGINT)))
		cmd.Process.Signal(syscall.SIGINT)
		return s.LogsWriter.Close()
	case err := <-cmdErrChan:
		out.Write([]byte(color.RedString("Error %v", err)))
		return err
	case <-doneChan:
		out.Write([]byte(color.GreenString("Command done")))
		return s.LogsWriter.Close()
	}

}
