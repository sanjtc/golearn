package xcrawler

import (
	"net/http"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"gotest.tools/assert"
)

func TestFilterHTML(t *testing.T) {
	aNode := &html.Node{
		Type:      html.DocumentNode,
		DataAtom:  atom.A,
		Data:      "a",
		Namespace: "testNamespace",
		Attr: []html.Attribute{
			{Namespace: "testNamespace", Key: "testKey", Val: "testVal"},
		},
	}

	bNode := &html.Node{
		Type:      html.DocumentNode,
		DataAtom:  atom.A,
		Data:      "b",
		Namespace: "testNamespace",
		Attr: []html.Attribute{
			{Namespace: "testNamespace", Key: "testKey", Val: "testVal"},
		},
	}

	req := &request{}

	aElement := &htmlElement{node: aNode, req: req}
	bElement := &htmlElement{node: bNode, req: req}

	aNodeFilter := func(t HTMLElement) bool {
		return t.GetData() == "a"
	}

	type TestCase struct {
		e        HTMLElement
		filters  []HTMLFilter
		expected bool
	}

	testCases := []TestCase{
		{aElement, []HTMLFilter{aNodeFilter}, true},
		{bElement, []HTMLFilter{aNodeFilter}, false},
	}

	for _, testCase := range testCases {
		got := FilterHTML(testCase.e, testCase.filters...)
		assert.Equal(t, testCase.expected, got)
	}
}

func TestFilterRequest(t *testing.T) {
	rawReq := &http.Request{}
	c := &crawler{maxDepth: 1}

	unvaildReq := &request{rawReq: nil, depth: 1, c: c}
	vaildReq := &request{rawReq: rawReq, depth: 1, c: c}

	vaildReqFilter := func(req Request) bool {
		return req.GetRawReq() != nil
	}

	type TestCase struct {
		req      Request
		filters  []RequestFilter
		expected bool
	}

	testCases := []TestCase{
		{unvaildReq, []RequestFilter{vaildReqFilter}, false},
		{vaildReq, []RequestFilter{vaildReqFilter}, true},
	}

	for _, testCase := range testCases {
		got := FilterRequest(testCase.req, testCase.filters...)
		assert.Equal(t, testCase.expected, got)
	}
}

func TestFilterResponse(t *testing.T) {
	errorRawResp := &http.Response{StatusCode: 401}
	errorResp := &response{rawResp: errorRawResp}

	successRawResp := &http.Response{StatusCode: 200}
	successResp := &response{rawResp: successRawResp}

	successRespFilter := func(resp Response) bool {
		return resp.GetStatusCode() < 400
	}

	type TestCase struct {
		resp     Response
		filters  []ResponseFilter
		expected bool
	}

	testCases := []TestCase{
		{errorResp, []ResponseFilter{successRespFilter}, false},
		{successResp, []ResponseFilter{successRespFilter}, true},
	}

	for _, testCase := range testCases {
		got := FilterResponse(testCase.resp, testCase.filters...)
		assert.Equal(t, testCase.expected, got)
	}
}
