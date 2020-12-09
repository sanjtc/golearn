package crawlerutil

import (
	"net/url"
	"strings"
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

func filterHrefWithAbsolutePath(u *url.URL) bool {
	return u.Scheme == "https" || u.Scheme == "http"
}

func filterHrefWithRelativePath(u *url.URL) bool {
	return u.Scheme == ""
}

func filterHrefWithFile(u *url.URL) bool {
	i := strings.Index(u.Path, ".")

	return i != -1
}

func filterHrefWithoutFile(u *url.URL) bool {
	i := strings.Index(u.Path, ".")

	return i == -1
}

func filterHrefWithJS(u *url.URL) bool {
	return u.Scheme == "javascript"
}

func filterHrefWithQuery(u *url.URL) bool {
	return u.RawQuery != ""
}

func filterHrefWithoutQuery(u *url.URL) bool {
	return u.RawQuery == ""
}
