package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nigimaxx/procgo/client/pkg"
	"github.com/nigimaxx/procgo/proto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func init() {
	startCmd.PersistentFlags().BoolP("detach", "d", false, "detach")
}

var startCmd = &cobra.Command{
	Use:               "start [services ...]",
	Short:             "start",
	Long:              `start`,
	PersistentPreRunE: pkg.CreateConnectPreRun(procfile, setClient, pkg.WithStartDaemon),
	RunE: func(cmd *cobra.Command, args []string) error {
		flagSet := cmd.Flags()
		detach, _ := flagSet.GetBool("detach")

		services, err := pkg.ParseAndSelect(procfile, args)
		if err != nil {
			return err
		}

		if _, err := client.Start(context.Background(), &proto.Services{Services: services}); err != nil {
			return err
		}

		// skip logs if -d
		if detach {
			return nil
		}

		stream, err := client.Logs(context.Background(), &proto.AllOrServices{All: true})
		if err != nil {
			return err
		}

		errChan := make(chan error)
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sigChan

			if err := stream.CloseSend(); err != nil {
				errChan <- errors.Wrap(err, "close send")
			}

			if _, err := client.KillAll(context.Background(), &emptypb.Empty{}); err != nil {
				errChan <- errors.Wrap(err, "kill all")
			}
		}()

		go func() {
			for {
				logLine, err := stream.Recv()
				if err == io.EOF {
					errChan <- nil
					break
				}

				if err != nil {
					st, _ := status.FromError(err)
					log.Println(st.Code(), st.Message(), st.Details())

					errChan <- errors.Wrap(err, "logs")
					break
				}

				fmt.Print(string(logLine.Value))
			}

			errChan <- nil
		}()

		return <-errChan
	},
}
