package pkg

import (
	"io/ioutil"
	"strings"

	"github.com/nigimaxx/procgo/proto"
)

func ParseProcfile(procfile string) ([]*proto.ServiceDefinition, error) {
	data, err := ioutil.ReadFile(procfile)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	services := []*proto.ServiceDefinition{}

	for _, line := range lines {
		split := strings.Split(line, ":")

		if strings.HasPrefix(line, "#") || len(split) < 2 || len(split[0]) == 0 || len(split[1]) == 0 {
			continue
		}

		services = append(services, &proto.ServiceDefinition{
			Name:    strings.TrimSpace(split[0]),
			Command: strings.TrimSpace(strings.Join(split[1:], ":")),
		})
	}

	return services, nil
}

// ParseAndSelect parse the procfile and returns the services mentioned in the args or all
func ParseAndSelect(procfile string, args []string) ([]*proto.ServiceDefinition, error) {
	allServices, err := ParseProcfile(procfile)
	if err != nil {
		return nil, err
	}

	services := []*proto.ServiceDefinition{}

	for _, s := range allServices {
		if len(args) == 0 || InStringList(args, s.Name) {
			services = append(services, s)
		}
	}

	return services, nil
}

func InStringList(list []string, item string) bool {
	for _, i := range list {
		if item == i {
			return true
		}
	}
	return false
}
