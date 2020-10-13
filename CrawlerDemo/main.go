package main

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/pantskun/golearn/CrawlerDemo/crawler"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
)

func main() {
	url := "https://www.ssetech.com.cn/"
	nodes := crawler.GetElementNodesFromURL(url, "a")

	urls := []string{}

	for _, n := range nodes {
		url := crawler.GetElementAttributeValue(n, "href")
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

	urls = crawler.FilterURL(urls, urlPrefixFilter, urlHTMLFilter)

	for _, url := range urls {
		log.Println(url)
	}

	interactor := etcd.NewInteractor()
	defer interactor.Close()

	for _, url := range urls {
		res, err := interactor.Get(url)
		if err != nil {
			log.Println(err)
			continue
		}

		if res == "" {
			err := interactor.Put(url, "1")
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println(url)
		}
	}
}
