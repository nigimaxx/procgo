package cmd

import (
	"context"

	"github.com/nigimaxx/procgo/client/pkg"
	"github.com/nigimaxx/procgo/proto"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [services ...]",
	Short: "restart",
	Long:  `restart`,
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
