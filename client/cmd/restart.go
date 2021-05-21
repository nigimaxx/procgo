package cmd

import (
	"context"

	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:     "restart [services ...]",
	Short:   "restart",
	Long:    `restart`,
	PreRunE: createConnectPreRun(),
	RunE: func(cmd *cobra.Command, args []string) error {
		allServices, err := pkg.ParseProcfile(procfile)
		if err != nil {
			return err
		}

		services := []*proto.ServiceDefinition{}

		for _, s := range allServices {
			if len(args) == 0 || contains(args, s.Name) {
				services = append(services, s)
			}
		}

		if _, err := client.Restart(context.Background(), &proto.Services{Services: services}); err != nil {
			return err
		}

		return nil
	},
}
