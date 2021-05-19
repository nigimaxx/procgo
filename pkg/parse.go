package pkg

import (
	"io/ioutil"
	"strings"
)

func ParseProcfile(procfile string) ([]Service, error) {
	data, err := ioutil.ReadFile(procfile)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	services := []Service{}

	for _, line := range lines {
		split := strings.Split(line, ":")

		if strings.HasPrefix(line, "#") || len(split) < 2 || len(split[0]) == 0 || len(split[1]) == 0 {
			continue
		}

		services = append(services, Service{
			Name:    strings.TrimSpace(split[0]),
			Command: strings.TrimSpace(strings.Join(split[1:], ":")),
		})
	}

	return services, nil
}
