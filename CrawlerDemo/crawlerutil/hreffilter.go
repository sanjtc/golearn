package crawlerutil

import (
	"net/url"
)

type HrefFilter func(u *url.URL) bool

func FilterHref(u *url.URL, filters ...HrefFilter) bool {
	for _, filter := range filters {
		if !filter(u) {
			return false
		}
	}

	return true
}

func filterHrefWithHTTP(u *url.URL) bool {
	return u.Scheme == "https" || u.Scheme == "http"
}

func filterHrefWithJS(u *url.URL) bool {
	return u.Scheme == "javascript"
}
