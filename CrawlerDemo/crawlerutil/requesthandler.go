package crawlerutil

import (
	"net/http"
	"path"

	"github.com/pantskun/golearn/CrawlerDemo/etcd"
)

func HandleRequestWithSync(req *http.Request, etcdInteractor etcd.Interactor) {
	key := generateKeyForCrawler(path.Join("request", req.URL.Host, req.URL.Path))
	if !Synchronize(key, etcdInteractor) {
		req.URL = nil
	}
}
