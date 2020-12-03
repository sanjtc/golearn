package htmlutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/pantskun/commonutils/pathutils"
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
func DownloadURL(url string, des string) error {
	rsp, err := http.Get(fmt.Sprint(url))
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	filePath := path.Join(des, pathutils.GetURLPath(url))

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	return writeToFile(filePath, body)
}
