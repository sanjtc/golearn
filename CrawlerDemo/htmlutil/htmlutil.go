package htmlutil

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/PuerkitoBio/goquery"
	"github.com/pantskun/commonutils/pathutils"
	"golang.org/x/net/html"
)

// GetElementNodesFromURL
// 从url页面获取element类型的标签
func GetElementNodesFromURL(url string, element string) []*html.Node {
	resp, err := http.Get(fmt.Sprint(url))
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Println(err)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
	}

	return doc.Find("a").Nodes
}

// GetElementAttributeValue
// 从element中获取attribute属性值
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

func FilterURL(url string, filters ...URLFilter) bool {
	for _, filter := range filters {
		if !filter(url) {
			return false
		}
	}

	return true
}

// FilterURLs
// 根据filters对urls进行过滤
func FilterURLs(urls []string, filters ...URLFilter) []string {
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

// DownloadURL
// 下载url至p位置
func DownloadURL(url string, p string) error {
	rsp, err := http.Get(fmt.Sprint(url))
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	filePath := path.Join(p, pathutils.GetURLPath(url))

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	return WriteToFile(filePath, body)
}
