package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/nigimaxx/procgo/daemon/handler"
	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
	"google.golang.org/grpc"
)

func main() {
	color.NoColor = false

	f, err := os.OpenFile("procgo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	conn, err := net.Listen("unix", pkg.SocketPath)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := handler.NewProcgoServer()

	grpcServer := grpc.NewServer(pkg.WithServerUnaryInterceptor(procfile))
	proto.RegisterProcgoServer(grpcServer, &server)

	doneChan := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-server.ErrChan:
				log.Println(err)
				close(server.KillChan)
				time.Sleep(1 * time.Second)
				close(doneChan)
			case <-server.DoneChan:
				log.Println("Done")
				if len(server.Services) == 0 {
					time.Sleep(1 * time.Second)
					close(doneChan)
				}
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

	log.Println("Stopping")
	grpcServer.Stop()
}
