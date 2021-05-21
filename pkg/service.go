package pkg

import (
	"os"
	"os/exec"
	"sync"
	"syscall"

	"github.com/fatih/color"
	"github.com/nigimaxx/procgo/proto"
)

var (
	shell      string
	colorIndex = -1
	colors     = []color.Attribute{
		color.FgRed,
		color.FgGreen,
		color.FgYellow,
		color.FgBlue,
		color.FgMagenta,
		color.FgCyan,
		color.FgWhite,
	}
)

func init() {
	shell = os.Getenv("SHELL")
	if shell == "" {
		panic("$SHELL not defined")
	}
}

type Service struct {
	Name     string
	Command  string
	StopChan chan os.Signal

	colorName color.Attribute
	logChan   chan []byte
	listeners map[chan []byte]struct{}
	mu        sync.Mutex
}

func NewServiceFromDef(s *proto.ServiceDefinition) *Service {
	colorIndex = (colorIndex + 1) % len(colors)
	return &Service{
		Name:      s.Name,
		Command:   s.Command,
		StopChan:  make(chan os.Signal),
		colorName: colors[colorIndex],
		logChan:   make(chan []byte),
		listeners: make(map[chan []byte]struct{}),
	}
}

func (s *Service) Clone() *Service {
	return &Service{
		Name:      s.Name,
		Command:   s.Command,
		StopChan:  make(chan os.Signal),
		colorName: s.colorName,
		logChan:   make(chan []byte),
		listeners: make(map[chan []byte]struct{}),
	}
}

func (s *Service) AddListener(ch chan []byte) {
	s.mu.Lock()
	s.listeners[ch] = struct{}{}
	s.mu.Unlock()
}

func (s *Service) RemoveListener(ch chan []byte) {
	s.mu.Lock()
	delete(s.listeners, ch)
	s.mu.Unlock()
}

func (s *Service) Start(killChan <-chan os.Signal) error {
	cmd := exec.Command(shell, "-c", s.Command)

	out := PrefixWriter{s.Name, s.colorName, s.logChan}
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

	go func() {
		for {
			line, ok := <-s.logChan
			if !ok {
				for l := range s.listeners {
					close(l)
				}

				s.mu.Lock()
				s.listeners = make(map[chan []byte]struct{})
				s.mu.Unlock()
				break
			}

			s.mu.Lock()
			for l := range s.listeners {
				l <- line
			}
			s.mu.Unlock()
		}
	}()

	select {
	case <-killChan:
		out.Write([]byte(color.YellowString("Killing with signal %v", syscall.SIGINT)))
		cmd.Process.Signal(syscall.SIGINT)
		close(s.logChan)
		return nil
	case <-s.StopChan:
		out.Write([]byte(color.YellowString("Killing with signal %v", syscall.SIGINT)))
		cmd.Process.Signal(syscall.SIGINT)
		close(s.logChan)
		return nil
	case err := <-cmdErrChan:
		out.Write([]byte(color.RedString("Error %v", err)))
		close(s.logChan)
		return err
	case <-doneChan:
		out.Write([]byte(color.GreenString("Command done")))
		close(s.logChan)
		return nil
	}
}
