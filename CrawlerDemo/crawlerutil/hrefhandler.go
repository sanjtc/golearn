package crawlerutil

import (
	"log"
	"net/url"
)

type HrefHandler func(u *url.URL) bool

func NewHandlerWithFilters(handler HrefHandler, filters ...HrefFilter) HrefHandler {
	return func(u *url.URL) bool {
		if !FilterHref(u, filters...) {
			return false
		}

		return handler(u)
	}
}

func handleHrefWithHTTP(u *url.URL) bool {
	log.Println("HandleHrefWithHTTP:", u.String())

	return true
}

func handleHrefWithJS(u *url.URL) bool {
	log.Println("HandleHrefWithJS:", u.String())

	return false
}
