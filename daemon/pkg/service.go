package pkg

import (
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"

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
		color.FgHiRed,
		color.FgHiGreen,
		color.FgHiYellow,
		color.FgHiBlue,
		color.FgHiMagenta,
		color.FgHiCyan,
		color.FgHiWhite,
	}
)

func init() {
	shell = os.Getenv("SHELL")
	if shell == "" {
		panic("$SHELL not defined")
	}
}

func nextColor() color.Attribute {
	colorIndex = (colorIndex + 1) % len(colors)
	return colors[colorIndex]
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
	return &Service{
		Name:      s.Name,
		Command:   s.Command,
		StopChan:  make(chan os.Signal),
		colorName: nextColor(),
		logChan:   make(chan []byte),
		listeners: make(map[chan []byte]struct{}),
	}
}

func (s *Service) Clone() *Service {
	return &Service{
		Name:      s.Name,
		Command:   s.Command,
		StopChan:  make(chan os.Signal),
		colorName: nextColor(),
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

	// use as done channel of command as well if error is nil
	cmdErrChan := make(chan error)
	shutdownChan := make(chan struct{})

	go func() {
		cmdErrChan <- cmd.Wait()
	}()

	go func() {
		<-shutdownChan
		time.Sleep(250 * time.Millisecond) // 250ms grace period for logs

		for l := range s.listeners {
			close(l)
		}

		s.mu.Lock()
		s.listeners = make(map[chan []byte]struct{})
		s.mu.Unlock()
	}()

	go func() {
		for {
			line := <-s.logChan

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
		shutdownChan <- struct{}{}
		return nil

	case <-s.StopChan:
		out.Write([]byte(color.YellowString("Killing with signal %v", syscall.SIGINT)))
		cmd.Process.Signal(syscall.SIGINT)
		shutdownChan <- struct{}{}
		return nil

	case err := <-cmdErrChan:
		if err != nil {
			out.Write([]byte(color.RedString("Error %v", err)))
			shutdownChan <- struct{}{}
			return err
		}

		out.Write([]byte(color.GreenString("Command done")))
		shutdownChan <- struct{}{}
		return nil
	}
}
