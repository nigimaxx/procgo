package cmd

import (
	"context"

	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
	"github.com/spf13/cobra"
)

func contains(list []string, item string) bool {
	for _, i := range list {
		if item == i {
			return true
		}
	}
	return false
}

var startCmd = &cobra.Command{
	Use:   "start [services ...]",
	Short: "start",
	Long:  `start`,
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

		if _, err := client.Start(context.Background(), &proto.Services{Services: services}); err != nil {
			return err
		}

		// listen

		return nil
	},
}
