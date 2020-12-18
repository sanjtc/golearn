package crawlerutil

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pantskun/golearn/CrawlerDemo/mock/mock_xcrawler"
	"github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	"gotest.tools/assert"
)

func TestFilterElementWithURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	hrefElement := mock_xcrawler.NewMockHTMLElement(ctrl)
	hrefElement.EXPECT().GetAttr("href").Return("testhref")

	srcElement := mock_xcrawler.NewMockHTMLElement(ctrl)
	srcElement.EXPECT().GetAttr("href").Return("testsrc")

	unvaildElement := mock_xcrawler.NewMockHTMLElement(ctrl)
	unvaildElement.EXPECT().GetAttr("src").Return("")
	unvaildElement.EXPECT().GetAttr("href").Return("")

	type TestCase struct {
		element  xcrawler.HTMLElement
		expected bool
	}

	testCases := []TestCase{
		{hrefElement, true},
		{srcElement, true},
		{unvaildElement, false},
	}

	for _, testCase := range testCases {
		got := FilterElementWithURL(testCase.element)
		assert.Equal(t, testCase.expected, got)
	}
}
