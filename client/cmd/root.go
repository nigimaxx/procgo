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
		Use:   "procgo",
		Short: "procgo is tool to run local services concurrently",
		Long: `procgo is tool to run local services concurrently.
It is similar to foreman. The main difference is
that it consists of a client daemon architecture
which allows it to start/stop/restart services independently.`,
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

	rootCmd.InitDefaultHelpCmd()
	rootCmd.InitDefaultHelpFlag()
	rootCmd.InitDefaultVersionFlag()

	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(killCmd)
	rootCmd.AddCommand(completionCmd)
}

func SetVersion(version string) {
	rootCmd.Version = version
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.New(color.FgRed).Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
