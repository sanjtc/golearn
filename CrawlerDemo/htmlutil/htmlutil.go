package htmlutil

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"golang.org/x/net/html"
)

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

// DownloadURL
// 下载url至p位置
func DownloadURL(u *url.URL, des string) error {
	rsp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	filePath := path.Join(des, u.Host, u.Path)

	if !strings.Contains(u.Path, ".") {
		filePath = path.Join(filePath, "index.html")
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	return WriteToFile(filePath, body)
}

func GetDomainFromURL(url string) string {
	ss := strings.Split(url, "://")
	if len(ss) == 1 {
		return ""
	}

	s := ss[1]

	i := strings.Index(s, "/")
	if i == -1 {
		return s
	}

	return s[:i]
}
