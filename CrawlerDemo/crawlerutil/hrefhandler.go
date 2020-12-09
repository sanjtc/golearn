package crawlerutil

import (
	"log"
	"net/http"
	"net/url"

	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
)

type HrefHandler func(u *url.URL, c xcrawler.Crawler) bool

func NewHandlerWithFilters(handler HrefHandler, filters ...HrefFilter) HrefHandler {
	return func(u *url.URL, c xcrawler.Crawler) bool {
		if !FilterHref(u, filters...) {
			return false
		}

		return handler(u, c)
	}
}

func handleHrefWithAbsolutePath(u *url.URL, c xcrawler.Crawler) bool {
	log.Println("HandleHrefWithAbsolutePath:", u)

	return true
}

// func handleHrefWithAbsolutePathWithoutFileAndQuery(u *url.URL, c xcrawler.Crawler) bool {
// 	log.Println("HrefWithAbsolutePathAndNotFileWithoutParam:", u)

// 	u.Path = path.Join(u.Path, "index.html")

// 	return true
// }

func handleHrefWithRelativePath(u *url.URL, c xcrawler.Crawler) bool {
	log.Println("HandleHrefWithRelativePath:", u)

	host := c.GetHost()

	if host == "" {
		return false
	}

	u.Scheme = "https"
	u.Host = host

	return true
}

func handleHrefWithAbsolutePathWithQuery(u *url.URL, c xcrawler.Crawler) bool {
	log.Println("HrefWithAbsolutePathWithQuery:", u)

	return tryRedirectURL(u)
}

func handleHrefWithRelativePathWithQuery(u *url.URL, c xcrawler.Crawler) bool {
	log.Println("HrefWithRelativePathWithQuery", u)

	domain := c.GetHost()

	if domain == "" {
		return false
	}

	u.Scheme = "https"
	u.Host = domain

	return tryRedirectURL(u)
}

func handleHrefWithJS(u *url.URL, c xcrawler.Crawler) bool {
	log.Println("HrefWithJS:", u)

	return false
}

func tryRedirectURL(u *url.URL) bool {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(u.String())
	if err != nil {
		log.Println("error", err)
		return false
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther, http.StatusTemporaryRedirect:
		{
			newURL, err := (resp.Location())
			if err != nil {
				log.Println("error:", err)
				return false
			}

			*u = *newURL
			return true
		}
	}

	return false
}
