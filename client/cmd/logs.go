package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/nigimaxx/procgo/client/pkg"
	"github.com/nigimaxx/procgo/proto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:               "logs [services ...]",
	Short:             "prints the logs of the services",
	Long:              `prints the logs of the provided services or of all services if none is provided`,
	PersistentPreRunE: pkg.CreateConnectPreRun(procfile, setClient, pkg.WithStartDaemon),
	RunE: func(cmd *cobra.Command, args []string) error {
		allServices, err := pkg.ParseProcfile(procfile)
		if err != nil {
			return err
		}

		services := []*proto.ServiceDefinition{}

		for _, s := range allServices {
			if pkg.InStringList(args, s.Name) {
				services = append(services, s)
			}
		}

		stream, err := client.Logs(context.Background(), &proto.AllOrServices{All: len(args) == 0, Services: services})
		if err != nil {
			return err
		}

		for {
			logLine, err := stream.Recv()
			if err == io.EOF {
				return nil
			}

			if err != nil {
				return errors.Wrap(err, "logs")
			}

			fmt.Print(string(logLine.Value))
		}

	},
}
