package crawlerutil

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pantskun/golearn/CrawlerDemo/mock/mock_etcd"
	"github.com/pantskun/golearn/CrawlerDemo/mock/mock_xcrawler"
	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyForCrawler(t *testing.T) {
	url := "test"
	expected := "/crawler-test"
	got := generateKeyForCrawler(url)
	assert.Equal(t, expected, got)
}

func TestGetURLString(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	element := mock_xcrawler.NewMockHTMLElement(ctrl)
	element.EXPECT().GetAttr("href").Return("").AnyTimes()
	element.EXPECT().GetAttr("src").Return("").AnyTimes()

	assert.Equal(t, "", getURLString(element))
}

// test abspath from href.
func TestHandleElementWithURL1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	element := mock_xcrawler.NewMockHTMLElement(ctrl)
	request := mock_xcrawler.NewMockRequest(ctrl)
	etcdInteractor := mock_etcd.NewMockInteractor(ctrl)

	u, _ := url.Parse("https://www.test.com")

	reqU, _ := url.Parse("https://www.test.com")
	rawReq := &http.Request{URL: reqU}

	element.EXPECT().GetAttr("href").Return(u.String()).AnyTimes()
	element.EXPECT().GetAttr("src").Return("").AnyTimes()
	element.EXPECT().GetRequest().Return(request).AnyTimes()

	request.EXPECT().GetDepth().Return(1).AnyTimes()
	request.EXPECT().GetRawReq().Return(rawReq).AnyTimes()
	request.EXPECT().Visit("https://www.test.com").AnyTimes()

	etcdInteractor.EXPECT().TxnSync(gomock.Any()).Return(true)

	HandleElementWithURL(element, etcdInteractor)
}

// test relpath from src.
func TestHandleElementWithURL2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	element := mock_xcrawler.NewMockHTMLElement(ctrl)
	request := mock_xcrawler.NewMockRequest(ctrl)
	etcdInteractor := mock_etcd.NewMockInteractor(ctrl)

	u, _ := url.Parse("/test")

	reqU, _ := url.Parse("https://www.test.com")
	rawReq := &http.Request{URL: reqU}

	element.EXPECT().GetAttr("href").Return("").AnyTimes()
	element.EXPECT().GetAttr("src").Return(u.String()).AnyTimes()
	element.EXPECT().GetRequest().Return(request).AnyTimes()

	request.EXPECT().GetDepth().Return(1).AnyTimes()
	request.EXPECT().GetRawReq().Return(rawReq).AnyTimes()
	request.EXPECT().Visit("https://www.test.com/test").AnyTimes()

	etcdInteractor.EXPECT().TxnSync(gomock.Any()).Return(true)

	HandleElementWithURL(element, etcdInteractor)
}

// test sync false
func TestHandleElementWithURL3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	element := mock_xcrawler.NewMockHTMLElement(ctrl)
	request := mock_xcrawler.NewMockRequest(ctrl)
	etcdInteractor := mock_etcd.NewMockInteractor(ctrl)

	u, _ := url.Parse("/test")

	reqU, _ := url.Parse("https://www.test.com")
	rawReq := &http.Request{URL: reqU}

	element.EXPECT().GetAttr("href").Return("").AnyTimes()
	element.EXPECT().GetAttr("src").Return(u.String()).AnyTimes()
	element.EXPECT().GetRequest().Return(request).AnyTimes()

	request.EXPECT().GetDepth().Return(1).AnyTimes()
	request.EXPECT().GetRawReq().Return(rawReq).AnyTimes()
	request.EXPECT().Visit("https://www.test.com/test").AnyTimes()

	etcdInteractor.EXPECT().TxnSync(gomock.Any()).Return(false)

	HandleElementWithURL(element, etcdInteractor)
}
