package htmlutil

import (
	"net/http"
	"net/url"
	"os"
	"path"
	"testing"

	"github.com/pantskun/commonutils/pathutils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestGetElementAttributeValue(t *testing.T) {
	type TestCase struct {
		element   *html.Node
		attribute string
		expected  string
	}

	element1 := &html.Node{
		Attr: []html.Attribute{
			{Namespace: "", Key: "key", Val: "value"},
		},
	}

	testCases := []TestCase{
		{element: element1, attribute: "key", expected: "value"},
		{element: element1, attribute: "error", expected: ""},
		{element: nil, attribute: "", expected: ""},
	}

	for _, testCase := range testCases {
		value := GetElementAttributeValue(testCase.element, testCase.attribute)
		assert.Equal(t, value, testCase.expected)
	}
}

func TestDownloadURL(t *testing.T) {
	modulePath := pathutils.GetModulePath("CrawlerDemo")
	downloadPath := path.Join(modulePath, "crawler/temp")

	defer os.RemoveAll(downloadPath)

	u, _ := url.Parse("https://www.baidu.com/index.html")

	err := DownloadURL(u, downloadPath)
	assert.Nil(t, err)
}

func TestHead(t *testing.T) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, _ := client.Get("https://www.ssetech.com.cn/statics/")

	location := resp.Header.Get("Location")

	t.Log(location)
}
