package pkg

import "github.com/nigimaxx/procgo/proto"

func InServiceList(list []*Service, name string) bool {
	for _, i := range list {
		if name == i.Name {
			return true
		}
	}
	return false
}

func InServiceDefList(list []*proto.ServiceDefinition, name string) bool {
	for _, i := range list {
		if name == i.Name {
			return true
		}
	}
	return false
}
