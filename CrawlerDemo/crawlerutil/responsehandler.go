package crawlerutil

import (
	"net/http"
	"path"
	"strings"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/pantskun/golearn/CrawlerDemo/etcd"
	"github.com/pantskun/golearn/CrawlerDemo/htmlutil"
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"github.com/pantskun/golearn/CrawlerDemo/xlogutil"
)

// HandleDownloadResp
// downloadURL中，若u发生重定向，则u更新为重定向后的地址.
func HandleDownloadResp(resp xcrawler.Response, etcdInteractor etcd.Interactor) {
	u := resp.GetRequest().GetRawReq().URL

	if resp.GetStatusCode() >= http.StatusBadRequest {
		xlogutil.Warning(u.String(), " response status ", resp.GetStatusCode())
		return
	}

	urlPath := u.Path

	pathSuffixIndex := strings.LastIndex(urlPath, "/")
	if pathSuffixIndex != -1 {
		pathSuffix := urlPath[pathSuffixIndex:]
		if !strings.Contains(pathSuffix, ".") {
			urlPath = path.Join(urlPath, "index.html")
		}
	} else {
		urlPath = path.Join(urlPath, "index.html")
	}

	defer func() {
		// 如果不是html页面，不需要继续访问
		if !(strings.Contains(urlPath, ".html") || strings.Contains(urlPath, ".htm")) {
			resp.Abandon()
			xlogutil.Warning("not html file, not need visit")
		}
	}()

	filePath := path.Join(pathutils.GetModulePath("CrawlerDemo"), "download", u.Host, urlPath)

	key := generateKeyForCrawler("download", u.String())
	if !etcdInteractor.TxnSync(key) {
		xlogutil.Warning(u.String(), " has been downloaded")
		return
	}

	fileContent := resp.GetBody()
	if len(fileContent) == 0 {
		xlogutil.Warning("no content, not download")
		return
	}

	if err := htmlutil.WriteToFile(filePath, fileContent); err != nil {
		xlogutil.Error(err)
		return
	}

	xlogutil.Warning("download: ", u.String(), " to ", filePath, "[response status: ", resp.GetStatusCode(), "]")
}
