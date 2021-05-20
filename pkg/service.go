package pkg

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/fatih/color"
	"github.com/nigimaxx/procgo/proto"
)

type Service struct {
	Name    string
	Command string
}

var shell string

func init() {
	shell = os.Getenv("SHELL")
	if shell == "" {
		panic("$SHELL not defined")
	}
}

func NewServiceFromDef(s *proto.ServiceDefinition) Service {
	return Service{s.Name, s.Command}
}

func (s *Service) ToDef() *proto.ServiceDefinition {
	return &proto.ServiceDefinition{Name: s.Name, Command: s.Command}
}

func (s *Service) Start(killChan <-chan os.Signal) error {
	cmd := exec.Command(shell, "-c", s.Command)

	out := NewPrefixOutput(s.Name)

	cmd.Stdout = out.Stdout
	cmd.Stderr = out.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	cmdErrChan := make(chan error)
	doneChan := make(chan struct{})

	go func() {
		if err := cmd.Wait(); err != nil {
			cmdErrChan <- err
		} else {
			doneChan <- struct{}{}
		}
	}()

	select {
	case <-killChan:
		out.Stdout.Write([]byte(color.YellowString("Killing with signal %v", syscall.SIGINT)))
		cmd.Process.Signal(syscall.SIGINT)
		return nil
	case err := <-cmdErrChan:
		out.Stdout.Write([]byte(color.RedString("Error %v", err)))
		return err
	case <-doneChan:
		out.Stdout.Write([]byte(color.GreenString("Command done")))
		return nil
	}

}
