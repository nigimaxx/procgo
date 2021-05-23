package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "kills the daemon",
	Long:  `kills the daemon if it isn't exiting`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.KillAll(context.Background(), &emptypb.Empty{})
		return err
	},
}
