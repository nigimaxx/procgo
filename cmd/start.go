package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nigimaxx/procgo/pkg"
	"github.com/spf13/cobra"
)

func contains(list []string, item string) bool {
	for _, i := range list {
		if item == i {
			return true
		}
	}
	return false
}

var startCmd = &cobra.Command{
	Use:   "start [services ...]",
	Short: "start",
	Long:  `start`,
	RunE: func(cmd *cobra.Command, args []string) error {
		allServices, err := pkg.ParseProcfile(procfile)
		if err != nil {
			return err
		}

		services := []pkg.Service{}

		for _, s := range allServices {
			if contains(args, s.Name) {
				services = append(services, s)
			}
		}

		errChan := make(chan error, len(services))
		killChan := make(chan os.Signal, len(services))
		sigChan := make(chan os.Signal, 1)
		doneChan := make(chan struct{}, 1)

		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		for _, s := range services {
			go func(s pkg.Service) {
				if err := s.Start(killChan); err != nil {
					errChan <- err
				}
				doneChan <- struct{}{}
			}(s)
		}

		go func() {
			signal := <-sigChan
			for range services {
				killChan <- signal
			}
		}()

		doneCount := 0

		for {
			select {
			case err := <-errChan:
				killChan <- syscall.SIGINT
				time.Sleep(1 * time.Second)
				return err

			case <-doneChan: // wait for all dones
				doneCount++
				if doneCount >= len(services) {
					return nil
				}
			}
		}

	},
}
