package crawlerutil

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pantskun/golearn/CrawlerDemo/mock/mock_etcd"
	"github.com/pantskun/golearn/CrawlerDemo/mock/mock_xcrawler"
)

func TestHandleDownloadResp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := mock_xcrawler.NewMockResponse(ctrl)
	etcdInteractor := mock_etcd.NewMockInteractor(ctrl)
	req := mock_xcrawler.NewMockRequest(ctrl)

	u, _ := url.Parse("www.test.com")
	rawReq := &http.Request{URL: u}

	req.EXPECT().GetRawReq().Return(rawReq).AnyTimes()

	resp.EXPECT().GetRequest().Return(req).AnyTimes()
	resp.EXPECT().GetStatusCode().Return(200).AnyTimes()
	resp.EXPECT().Abandon().AnyTimes()
	resp.EXPECT().GetBody().Return([]byte{}).AnyTimes()

	etcdInteractor.EXPECT().TxnSync(gomock.Any()).Return(true).AnyTimes()

	HandleDownloadResp(resp, etcdInteractor)
}

// test url not html file and sync false.
func TestHandleDownloadResp2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := mock_xcrawler.NewMockResponse(ctrl)
	etcdInteractor := mock_etcd.NewMockInteractor(ctrl)
	req := mock_xcrawler.NewMockRequest(ctrl)

	u, _ := url.Parse("www.test.com/test.js")
	rawReq := &http.Request{URL: u}

	req.EXPECT().GetRawReq().Return(rawReq).AnyTimes()

	resp.EXPECT().GetRequest().Return(req).AnyTimes()
	resp.EXPECT().GetStatusCode().Return(200).AnyTimes()
	resp.EXPECT().Abandon().AnyTimes()
	resp.EXPECT().GetBody().Return([]byte{}).AnyTimes()

	etcdInteractor.EXPECT().TxnSync(gomock.Any()).Return(false).AnyTimes()

	HandleDownloadResp(resp, etcdInteractor)
}

// test url has path.
func TestHandleDownloadResp3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := mock_xcrawler.NewMockResponse(ctrl)
	etcdInteractor := mock_etcd.NewMockInteractor(ctrl)
	req := mock_xcrawler.NewMockRequest(ctrl)

	u, _ := url.Parse("www.test.com/test")
	rawReq := &http.Request{URL: u}

	req.EXPECT().GetRawReq().Return(rawReq).AnyTimes()

	resp.EXPECT().GetRequest().Return(req).AnyTimes()
	resp.EXPECT().GetStatusCode().Return(200).AnyTimes()
	resp.EXPECT().Abandon().AnyTimes()
	resp.EXPECT().GetBody().Return([]byte{}).AnyTimes()

	etcdInteractor.EXPECT().TxnSync(gomock.Any()).Return(true).AnyTimes()

	HandleDownloadResp(resp, etcdInteractor)
}

// test status code >= 400
func TestHandleDownloadResp4(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	resp := mock_xcrawler.NewMockResponse(ctrl)
	etcdInteractor := mock_etcd.NewMockInteractor(ctrl)
	req := mock_xcrawler.NewMockRequest(ctrl)

	u, _ := url.Parse("www.test.com/test")
	rawReq := &http.Request{URL: u}

	req.EXPECT().GetRawReq().Return(rawReq).AnyTimes()

	resp.EXPECT().GetRequest().Return(req).AnyTimes()
	resp.EXPECT().GetStatusCode().Return(401).AnyTimes()
	resp.EXPECT().Abandon().AnyTimes()
	resp.EXPECT().GetBody().Return([]byte{}).AnyTimes()

	etcdInteractor.EXPECT().TxnSync(gomock.Any()).Return(true).AnyTimes()

	HandleDownloadResp(resp, etcdInteractor)
}
