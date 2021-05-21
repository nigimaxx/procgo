package cmd

import (
	"github.com/nigimaxx/procgo/pkg"
	"github.com/nigimaxx/procgo/proto"
)

// parseAndSelect parse the procfile and returns the services mentioned in the args or all
func parseAndSelect(args []string) ([]*proto.ServiceDefinition, error) {
	allServices, err := pkg.ParseProcfile(procfile)
	if err != nil {
		return nil, err
	}

	services := []*proto.ServiceDefinition{}

	for _, s := range allServices {
		if len(args) == 0 || contains(args, s.Name) {
			services = append(services, s)
		}
	}

	return services, nil
}
