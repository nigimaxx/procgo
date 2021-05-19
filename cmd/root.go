package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	procfile string

	rootCmd = &cobra.Command{
		Use:           "procgo",
		Short:         "procgo",
		Long:          `procgo`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&procfile, "procfile", "j", "Procfile", "procfile")

	rootCmd.AddCommand(startCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}
