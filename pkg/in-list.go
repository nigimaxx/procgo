package pkg

func InServiceList(list []Service, name string) bool {
	for _, i := range list {
		if name == i.Name {
			return true
		}
	}
	return false
}
