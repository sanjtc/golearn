package crawlerutil

import (
	"net/url"
)

type URLFilter func(u *url.URL) bool

func FilterURL(u *url.URL, filters ...URLFilter) bool {
	for _, filter := range filters {
		if !filter(u) {
			return false
		}
	}

	return true
}

func filterURLWithHTTP(u *url.URL) bool {
	return u.Scheme == "https" || u.Scheme == "http"
}

func filterURLWithJS(u *url.URL) bool {
	return u.Scheme == "javascript"
}
