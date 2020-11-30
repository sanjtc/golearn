package main

import (
	"flag"
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
)

func main() {
	var url string

	flag.StringVar(&url, "url", "https://www.ssetech.com.cn/", "url")
	flag.Parse()

	// 获取url
	nodes := htmlutil.GetElementNodesFromURL(url, "a")

	urls := []string{}

	for _, n := range nodes {
		url := htmlutil.GetElementAttributeValue(n, "href")
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

	urls = htmlutil.FilterURL(urls, urlPrefixFilter, urlHTMLFilter)

	// 下载url
	// interactor := etcd.NewInteractorWithEmbed()
	etcdInteractor, err := etcd.NewInteractor()
	if err != nil {
		log.Println(err)
		return
	}
	defer etcdInteractor.Close()

	for _, url := range urls {
		// lock
		if _, err := etcdInteractor.Lock(); err != nil {
			log.Println(err)
			return
		}

		// check url
		res, err := etcdInteractor.Get(url)
		if err != nil {
			log.Println(err)
			return
		}

		needDownload := false

		if res == "" {
			err := etcdInteractor.Put(url, "1")
			if err != nil {
				log.Println(err)
				return
			}

			needDownload = true
		}
		// unlock
		if _, err := etcdInteractor.Unlock(); err != nil {
			log.Println(err)
			return
		}

		// download
		if needDownload {
			err = htmlutil.DownloadURL(url, path.Join(pathutils.GetModulePath("CrawlerDemo"), "download"))
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println(url)
		}
	}
}
