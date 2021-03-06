package cmd

import (
	"context"

	"github.com/nigimaxx/procgo/client/pkg"
	"github.com/nigimaxx/procgo/proto"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [services ...]",
	Short: "restarts the provided services",
	Long:  `restarts the provided services of all services if none is provided`,
	RunE: func(cmd *cobra.Command, args []string) error {
		services, err := pkg.ParseAndSelect(procfile, args)
		if err != nil {
			return err
		}

		if _, err := client.Restart(context.Background(), &proto.Services{Services: services}); err != nil {
			return err
		}

		return nil
	},
}
