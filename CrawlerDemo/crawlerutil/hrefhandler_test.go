package crawlerutil

import (
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestHrefHandler(t *testing.T) {
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
		handleHrefWithHTTP(u)
		got := u.String()
		assert.Equal(t, testCase.expected, got)
	}

	// handleHrefWithRelativePathAndHTMLFile
	// testRelativePathAndHTMLFileCase := []TestCase{
	// 	{"/test/test.html", "https://www.test.com/test/test.html"},
	// }
	// for _, testCase := range testRelativePathAndHTMLFileCase {
	// 	u, _ := url.Parse(testCase.href)
	// 	handleHrefWithRelativePath(u, c)
	// 	got := u.String()
	// 	assert.Equal(t, testCase.expected, got)
	// }

	// handleHrefWithJS
	testJSCase := []TestCase{
		{"javascript:void(0)", "javascript:void(0)"},
	}
	for _, testCase := range testJSCase {
		u, _ := url.Parse(testCase.href)
		handleHrefWithJS(u)
		got := u.String()
		assert.Equal(t, testCase.expected, got)
	}
}
