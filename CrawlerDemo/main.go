package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

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

	nodes := doc.Find("a").Nodes
	for _, n := range nodes {
		log.Println("Name: ", n.Data, " herf: ", GetElementAttributeValue(n, "href"))
	}
}
