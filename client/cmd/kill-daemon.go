package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"
)

var killDaemonCmd = &cobra.Command{
	Use:   "kill-daemon",
	Short: "kill-daemon",
	Long:  `kill-daemon`,
	// don't start if daemon down
	PreRunE: connectClient,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.KillAll(context.Background(), &emptypb.Empty{})
		return err
	},
}
