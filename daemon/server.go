//go:generate protoc --go_out=.. --go_opt=paths=source_relative --go-grpc_out=.. --go-grpc_opt=paths=source_relative --proto_path=.. ../proto/procgo.proto

package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/nigimaxx/procgo/common"
	"github.com/nigimaxx/procgo/daemon/handler"
	"github.com/nigimaxx/procgo/daemon/pkg"
	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/grpc"
)

func main() {
	color.NoColor = false

	f, err := os.OpenFile("/var/log/procgo/daemon.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatal("Missing procfile path")
	}

	procfile := os.Args[1]
	log.Println("Procfile", procfile)

	conn, err := net.Listen("unix", common.SocketPath)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := handler.NewProcgoServer()

	grpcServer := grpc.NewServer(pkg.WithServerUnaryInterceptor(procfile))
	proto.RegisterProcgoServer(grpcServer, &server)

	doneChan := make(chan struct{})

	go func() {
		for {
			err := <-server.ErrChan
			if err != nil {
				log.Println("Error:", err)

				close(server.KillChan)
				time.Sleep(1 * time.Second)
				close(doneChan)

				continue
			}

			if len(server.Services) == 0 {
				doneChan <- struct{}{}
			}
		}
	}()

	go func() {
		err := grpcServer.Serve(conn)
		if err != nil {
			log.Println(err)
			log.Fatal(err)
		}
	}()

	log.Println("Starting")

	<-doneChan

	log.Println("Stopping in 1s")
	time.Sleep(1 * time.Second)
	grpcServer.Stop()
}
