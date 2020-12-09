package crawlerutil

import (
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"gotest.tools/assert"
)

func TestHrefHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := xcrawler.NewMockCrawler(ctrl)
	c.EXPECT().GetHost().Return("www.test.com").Times(2)

	type TestCase struct {
		href     string
		expected string
	}

	// handleHrefWithAbsolutePathAndHTMLFile
	testAbsolutePathAndHTMLFileCases := []TestCase{
		{"https://test/test.html", "https://test/test.html"},
	}
	for _, testCase := range testAbsolutePathAndHTMLFileCases {
		u, _ := url.Parse(testCase.href)
		handleHrefWithAbsolutePath(u, c)
		got := u.String()
		assert.Equal(t, testCase.expected, got)
	}

	// handleHrefWithRelativePathAndHTMLFile
	testRelativePathAndHTMLFileCase := []TestCase{
		{"/test/test.html", "https://www.test.com/test/test.html"},
	}
	for _, testCase := range testRelativePathAndHTMLFileCase {
		u, _ := url.Parse(testCase.href)
		handleHrefWithRelativePath(u, c)
		got := u.String()
		assert.Equal(t, testCase.expected, got)
	}

	// handleHrefWithJS
	testJSCase := []TestCase{
		{"javascript:void(0)", "javascript:void(0)"},
	}
	for _, testCase := range testJSCase {
		u, _ := url.Parse(testCase.href)
		handleHrefWithJS(u, c)
		got := u.String()
		assert.Equal(t, testCase.expected, got)
	}
}
