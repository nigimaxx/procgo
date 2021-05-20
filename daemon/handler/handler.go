package handler

import (
	"os"

	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
)

type ProcgoServer struct {
	proto.UnimplementedProcgoServer
	Services []pkg.Service
	ErrChan  chan error
	KillChan chan os.Signal
	DoneChan chan struct{}
}

func NewProcgoServer() ProcgoServer {
	return ProcgoServer{
		Services: []pkg.Service{},
		ErrChan:  make(chan error),
		DoneChan: make(chan struct{}),
		KillChan: make(chan os.Signal, 1),
	}
}
