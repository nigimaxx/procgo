package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/nigimaxx/procgo/client/pkg"
	"github.com/nigimaxx/procgo/proto"
	"github.com/spf13/cobra"
)

var (
	procfile string
	client   proto.ProcgoClient

	rootCmd = &cobra.Command{
		Use:               "procgo",
		Short:             "procgo",
		Long:              `procgo`,
		SilenceErrors:     true,
		SilenceUsage:      true,
		PersistentPreRunE: pkg.CreateConnectPreRun(procfile, setClient),
	}
)

func setClient(c proto.ProcgoClient) {
	client = c
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&procfile, "procfile", "j", "Procfile", "procfile")

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(killCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}
