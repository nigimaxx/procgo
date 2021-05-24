package pkg

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/nigimaxx/procgo/common"
	"github.com/nigimaxx/procgo/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type options struct {
	start bool
}

type connectOption func(options) options

func WithStartDaemon(opts options) options {
	opts.start = true
	return opts
}

type setClient func(proto.ProcgoClient)

func CreateConnectPreRun(procfile string, setClient setClient, opts ...connectOption) func(*cobra.Command, []string) error {
	connOpts := options{}
	for _, o := range opts {
		connOpts = o(connOpts)
	}

	return func(_ *cobra.Command, _ []string) error {
		conn, err := grpc.Dial("unix://"+common.SocketPath, grpc.WithInsecure(), WithClientUnaryInterceptor(procfile))
		if err != nil {
			return err
		}

		client := proto.NewProcgoClient(conn)
		setClient(client)

		if !connOpts.start {
			return nil
		}

		if _, err := client.Ping(context.Background(), &emptypb.Empty{}); err != nil {
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.Unavailable {
				return startDaemon(procfile, client)
			}

			return err
		}

		return nil
	}
}

func startDaemon(procfile string, client proto.ProcgoClient) error {
	log.Println("Starting daemon")

	if err := os.RemoveAll(common.SocketPath); err != nil {
		return err
	}

	absProcfile, err := filepath.Abs(procfile)
	if err != nil {
		return err
	}

	cmd := exec.Command("procgo-daemon", absProcfile)
	// start process in its own process group to prevent ctrl+c to kill it
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Pgid: 0}
	if err := cmd.Start(); err != nil {
		return err
	}

	for {
		time.Sleep(100 * time.Millisecond)

		if _, err := client.Ping(context.Background(), &emptypb.Empty{}); err != nil {
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.Unavailable {
				continue
			}
			return err
		}

		break
	}

	return nil
}
