package crawlerutil

import (
	"path"
	"strings"
)

type HrefFilter func(string) bool

func FilterHref(href string, filters ...HrefFilter) bool {
	for _, filter := range filters {
		if !filter(href) {
			return false
		}
	}

	return true
}

func filterHrefWithAbsolutePath(href string) bool {
	return strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://")
}

func filterHrefWithRelativePath(href string) bool {
	return strings.HasPrefix(href, "/")
}

func filterHrefWithHTMLFile(href string) bool {
	ext := path.Ext(href)
	return ext == ".html"
}

func filterHrefWithNotFile(href string) bool {
	return strings.HasSuffix(href, "/")
}

func filterHrefWithJS(href string) bool {
	return strings.HasPrefix(href, "javascript:")
}
