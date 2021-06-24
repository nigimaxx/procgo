package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "lists all running services",
	Long:  `lists all running services with name and command`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := client.List(context.Background(), &emptypb.Empty{})
		if err != nil {
			return err
		}

		for _, svc := range s.Services {
			fmt.Printf("%-20s %s\n", svc.Name+":", svc.Command)
		}

		return nil
	},
}
