package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}

	log.Println("--------------get url start--------------")

	nodes := doc.Find("a").Nodes
	urls := []string{}

	for _, n := range nodes {
		url := GetElementAttributeValue(n, "href")
		log.Println("Name: ", n.Data, " herf: ", url)
		urls = append(urls, url)
	}

	log.Println("--------------get url end--------------")
	log.Println("--------------filter url start--------------")

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

	log.Println("--------------filter url end--------------")
}
