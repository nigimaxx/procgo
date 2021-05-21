package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"
)

var killCmd = &cobra.Command{
	Use:     "kill",
	Short:   "kill",
	Long:    `kill`,
	PreRunE: createConnectPreRun(),
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := client.KillAll(context.Background(), &emptypb.Empty{})
		return err
	},
}
