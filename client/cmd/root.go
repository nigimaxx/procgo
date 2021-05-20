package cmd

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	procfile string
	client   proto.ProcgoClient

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

	conn, err := grpc.Dial("unix://"+pkg.SocketPath, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client = proto.NewProcgoClient(conn)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}
