package crawlerutil

import (
	"path"
	"strings"
)

type URLFilter func(url string) bool

func filterURL(url string, filters ...URLFilter) bool {
	for _, filter := range filters {
		if !filter(url) {
			return false
		}
	}

	return true
}

func filterURLWithAbsolutePath(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func filterURLWithRelativePath(url string) bool {
	return strings.HasPrefix(url, "/")
}

func filterURLWithHTMLFile(url string) bool {
	ext := path.Ext(url)
	return ext == ".html"
}

func filterURLWithNotFile(url string) bool {
	return strings.HasSuffix(url, "/")
}

func filterURLWithJS(url string) bool {
	return strings.HasPrefix(url, "javascript:")
}
