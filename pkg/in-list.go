package pkg

func InServiceList(list []Service, item Service) bool {
	for _, i := range list {
		if item.Name == i.Name {
			return true
		}
	}
	return false
}
