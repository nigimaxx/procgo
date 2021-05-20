package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/nigimaxx/procgo/daemon/handler"
	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/grpc"
)

func main() {
	// outfile, err := os.Create("./procgo.log")
	// if err != nil {
	// 	log.Fatalf("failed to open log file: %v", err)
	// }

	// os.Stdout = outfile

	if err := os.RemoveAll(pkg.SocketPath); err != nil {
		log.Fatalf("failed to remove socket: %v", err)
	}

	conn, err := net.Listen("unix", pkg.SocketPath)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := handler.NewProcgoServer()

	grpcServer := grpc.NewServer()
	proto.RegisterProcgoServer(grpcServer, &server)

	doneChan := make(chan struct{})

	go func() {
		for {
			select {
			case <-server.ErrChan:
				close(server.KillChan)
				time.Sleep(1 * time.Second)
				close(doneChan)
			case <-server.DoneChan:
				if len(server.Services) == 0 {
					close(doneChan)
				}
			}
		}
	}()

	go func() {
		err := grpcServer.Serve(conn)
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-doneChan
}
