package crawlerutil

type URLFilter func(url string) bool

func FilterURL(url string, filters ...URLFilter) bool {
	for _, filter := range filters {
		if !filter(url) {
			return false
		}
	}

	return true
}
