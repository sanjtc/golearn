package main

import (
	"log"
	"path"
	"strings"

	"github.com/pantskun/golearn/CrawlerDemo/crawler"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"golang.org/x/net/html"
)

// GetElementAttributeValue get attribute value from html.Node.
func GetElementAttributeValue(element *html.Node, attribute string) string {
	if element == nil {
		return ""
	}

	for _, attr := range element.Attr {
		if attr.Key == attribute {
			return attr.Val
		}
	}

	return ""
}

type URLFilter func(string) bool

func FilterURL(urls []string, filters ...URLFilter) []string {
	result := []string{}

	for _, url := range urls {
		need := true

		for _, filter := range filters {
			if !filter(url) {
				need = false
				break
			}
		}

		if need {
			result = append(result, url)
		}
	}

	return result
}

func main() {
	url := "https://www.ssetech.com.cn/"
	nodes := crawler.GetElementNodesFromURL(url, "a")

	urls := []string{}

	for _, n := range nodes {
		url := GetElementAttributeValue(n, "href")
		log.Println("Name: ", n.Data, " herf: ", url)
		urls = append(urls, url)
	}

	urlPrefixFilter := func(url string) bool {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			return true
		}

		return false
	}

	urlHTMLFilter := func(url string) bool {
		ext := path.Ext(url)
		return ext == ".html"
	}

	urls = FilterURL(urls, urlPrefixFilter, urlHTMLFilter)

	for _, url := range urls {
		log.Println(url)
	}

	interactor := etcd.NewInteractor()
}
