package pkg

import (
	"fmt"
	"os"
	"os/exec"
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
	case signal := <-killChan:
		out.Stdout.Write([]byte(fmt.Sprintf("Killing with signal %v", signal)))
		cmd.Process.Signal(signal)
		return nil
	case err := <-cmdErrChan:
		return err
	case <-doneChan:
		return nil
	}

}
