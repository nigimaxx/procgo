package cmd

import (
	"context"

	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop [services ...]",
	Short:   "stop",
	Long:    `stop`,
	PreRunE: connectClient,
	RunE: func(cmd *cobra.Command, args []string) error {
		allServices, err := pkg.ParseProcfile(procfile)
		if err != nil {
			return err
		}

		services := []*proto.ServiceDefinition{}

		for _, s := range allServices {
			if len(args) == 0 || contains(args, s.Name) {
				services = append(services, s.ToDef())
			}
		}

		if _, err := client.Stop(context.Background(), &proto.Services{Services: services}); err != nil {
			return err
		}

		return nil
	},
}
