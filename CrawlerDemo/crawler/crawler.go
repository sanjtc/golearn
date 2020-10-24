package crawler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/pantskun/commonutils/pathutils"
	"golang.org/x/net/html"
)

func GetElementNodesFromURL(url string, element string) []*html.Node {
	resp, err := http.Get(fmt.Sprint(url))
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

	return doc.Find("a").Nodes
}

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

// FilterURL filter url by URLFilters.
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

func DownloadURL(url string) error {
	rsp, err := http.Get(fmt.Sprint(url))
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	path := pathutils.GetModulePath("CrawlerDemo") + "/download/" + pathutils.GetURLPath(url)

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	file := HTMLFile{Path: path, Content: body}

	return file.Write()
}
