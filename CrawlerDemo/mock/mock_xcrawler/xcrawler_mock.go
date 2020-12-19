// Code generated by MockGen. DO NOT EDIT.
// Source: .\xcrawler\xcrawler.go

// Package mock_xcrawler is a generated GoMock package.
package mock_xcrawler

import (
	gomock "github.com/golang/mock/gomock"
	xcrawler "github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	url "net/url"
	reflect "reflect"
)

// MockCrawler is a mock of Crawler interface
type MockCrawler struct {
	ctrl     *gomock.Controller
	recorder *MockCrawlerMockRecorder
}

// MockCrawlerMockRecorder is the mock recorder for MockCrawler
type MockCrawlerMockRecorder struct {
	mock *MockCrawler
}

// NewMockCrawler creates a new mock instance
func NewMockCrawler(ctrl *gomock.Controller) *MockCrawler {
	mock := &MockCrawler{ctrl: ctrl}
	mock.recorder = &MockCrawlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrawler) EXPECT() *MockCrawlerMockRecorder {
	return m.recorder
}

// Visit mocks base method
func (m *MockCrawler) Visit(url string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Visit", url)
}

// Visit indicates an expected call of Visit
func (mr *MockCrawlerMockRecorder) Visit(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Visit", reflect.TypeOf((*MockCrawler)(nil).Visit), url)
}

// AddHTMLHandler mocks base method
func (m *MockCrawler) AddHTMLHandler(handler xcrawler.HTMLHandler, filters ...xcrawler.HTMLFilter) {
	m.ctrl.T.Helper()
	varargs := []interface{}{handler}
	for _, a := range filters {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddHTMLHandler", varargs...)
}

// AddHTMLHandler indicates an expected call of AddHTMLHandler
func (mr *MockCrawlerMockRecorder) AddHTMLHandler(handler interface{}, filters ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{handler}, filters...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHTMLHandler", reflect.TypeOf((*MockCrawler)(nil).AddHTMLHandler), varargs...)
}

// AddRequestHandler mocks base method
func (m *MockCrawler) AddRequestHandler(handler xcrawler.RequestHandler, filters ...xcrawler.RequestFilter) {
	m.ctrl.T.Helper()
	varargs := []interface{}{handler}
	for _, a := range filters {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddRequestHandler", varargs...)
}

// AddRequestHandler indicates an expected call of AddRequestHandler
func (mr *MockCrawlerMockRecorder) AddRequestHandler(handler interface{}, filters ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{handler}, filters...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRequestHandler", reflect.TypeOf((*MockCrawler)(nil).AddRequestHandler), varargs...)
}

// AddResponseHandler mocks base method
func (m *MockCrawler) AddResponseHandler(handler xcrawler.ResponseHandler, filters ...xcrawler.ResponseFilter) {
	m.ctrl.T.Helper()
	varargs := []interface{}{handler}
	for _, a := range filters {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddResponseHandler", varargs...)
}

// AddResponseHandler indicates an expected call of AddResponseHandler
func (mr *MockCrawlerMockRecorder) AddResponseHandler(handler interface{}, filters ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{handler}, filters...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResponseHandler", reflect.TypeOf((*MockCrawler)(nil).AddResponseHandler), varargs...)
}

// visit mocks base method
func (m *MockCrawler) visit(u *url.URL, depth int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "visit", u, depth)
}

// visit indicates an expected call of visit
func (mr *MockCrawlerMockRecorder) visit(u, depth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "visit", reflect.TypeOf((*MockCrawler)(nil).visit), u, depth)
}