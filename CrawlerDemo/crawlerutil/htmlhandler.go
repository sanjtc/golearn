package crawlerutil

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

func HandleElementWithURL(element xcrawler.HTMLElement, etcdInteractor etcd.Interactor) {
	var s string

	href := element.GetAttr("href")
	src := element.GetAttr("src")

	switch {
	case href != "":
		{
			s = href
		}
	case src != "":
		{
			s = src
		}
	default:
		{
			return
		}
	}

	elementKey := generateKeyForCrawler(path.Join("element", element.GetOwner().Host, element.GetOwner().Path, s))
	if !Synchronize(elementKey, etcdInteractor) {
		return
	}

	log.Println("process:", s, " owner:", element.GetOwner())

	u, err := url.Parse(s)
	if err != nil {
		xlogutil.Error(err)
		return
	}

	if u.Host == "" {
		u.Host = element.GetOwner().Host
	}

	if u.Scheme == "" {
		u.Scheme = element.GetOwner().Scheme
	}

	hrefHandlers := []HrefHandler{
		NewHandlerWithFilters(handleHrefWithHTTP, filterHrefWithHTTP),
		NewHandlerWithFilters(handleHrefWithJS, filterHrefWithJS),
	}

	for _, handler := range hrefHandlers {
		if handler(u) {
			downloadURL(u, etcdInteractor)

			if strings.HasSuffix(u.Path, ".html") || !strings.Contains(u.Path, ".") {
				element.Visit(u)
			}
		}
	}
}

func generateKeyForCrawler(url string) string {
	return path.Join("/crawler", url)
}

func downloadURL(u *url.URL, etcdInteractor etcd.Interactor) {
	if u == nil {
		return
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(u.String())
	if err != nil {
		log.Println("error", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		log.Println("warning:", u.String(), " response status ", resp.Status)
		return
	}

	fileContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		xlogutil.Error(err)
		return
	}

	urlHost := u.Host
	urlPath := u.Path

	if redirectURL, _ := resp.Location(); redirectURL != nil {
		urlHost = redirectURL.Host
		urlPath = redirectURL.Path
	}

	if !strings.Contains(urlPath, ".") {
		urlPath = path.Join(urlPath, "index.html")
	}

	filePath := path.Join(pathutils.GetModulePath("CrawlerDemo"), "download", urlHost, urlPath)

	key := generateKeyForCrawler(path.Join("/download", urlHost, urlPath))
	if !Synchronize(key, etcdInteractor) {
		return
	}

	if err := htmlutil.WriteToFile(filePath, fileContent); err != nil {
		xlogutil.Error(err)
		return
	}

	log.Println("download:", u)
}
