package cmd

import (
	"context"

	"github.com/nigimaxx/procgo/proto"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:     "stop [services ...]",
	Short:   "stop",
	Long:    `stop`,
	PreRunE: createConnectPreRun(),
	RunE: func(cmd *cobra.Command, args []string) error {
		services, err := parseAndSelect(args)
		if err != nil {
			return err
		}

		if _, err := client.Stop(context.Background(), &proto.Services{Services: services}); err != nil {
			return err
		}

		return nil
	},
}
