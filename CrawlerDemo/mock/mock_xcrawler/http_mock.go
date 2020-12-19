// Code generated by MockGen. DO NOT EDIT.
// Source: .\xcrawler\http.go

// Package mock_xcrawler is a generated GoMock package.
package mock_xcrawler

import (
	gomock "github.com/golang/mock/gomock"
	xcrawler "github.com/pantskun/golearn/CrawlerDemo/xcrawler"
	http "net/http"
	reflect "reflect"
)

// MockRequest is a mock of Request interface
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// IsValid mocks base method
func (m *MockRequest) IsValid() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValid indicates an expected call of IsValid
func (mr *MockRequestMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockRequest)(nil).IsValid))
}

// GetDepth mocks base method
func (m *MockRequest) GetDepth() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDepth")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetDepth indicates an expected call of GetDepth
func (mr *MockRequestMockRecorder) GetDepth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDepth", reflect.TypeOf((*MockRequest)(nil).GetDepth))
}

// GetRawReq mocks base method
func (m *MockRequest) GetRawReq() *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawReq")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// GetRawReq indicates an expected call of GetRawReq
func (mr *MockRequestMockRecorder) GetRawReq() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawReq", reflect.TypeOf((*MockRequest)(nil).GetRawReq))
}

// Visit mocks base method
func (m *MockRequest) Visit(URL string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Visit", URL)
}

// Visit indicates an expected call of Visit
func (mr *MockRequestMockRecorder) Visit(URL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Visit", reflect.TypeOf((*MockRequest)(nil).Visit), URL)
}

// MockResponse is a mock of Response interface
type MockResponse struct {
	ctrl     *gomock.Controller
	recorder *MockResponseMockRecorder
}

// MockResponseMockRecorder is the mock recorder for MockResponse
type MockResponseMockRecorder struct {
	mock *MockResponse
}

// NewMockResponse creates a new mock instance
func NewMockResponse(ctrl *gomock.Controller) *MockResponse {
	mock := &MockResponse{ctrl: ctrl}
	mock.recorder = &MockResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResponse) EXPECT() *MockResponseMockRecorder {
	return m.recorder
}

// GetStatusCode mocks base method
func (m *MockResponse) GetStatusCode() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatusCode")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetStatusCode indicates an expected call of GetStatusCode
func (mr *MockResponseMockRecorder) GetStatusCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatusCode", reflect.TypeOf((*MockResponse)(nil).GetStatusCode))
}

// GetStatus mocks base method
func (m *MockResponse) GetStatus() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockResponseMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockResponse)(nil).GetStatus))
}

// GetRequest mocks base method
func (m *MockResponse) GetRequest() xcrawler.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest")
	ret0, _ := ret[0].(xcrawler.Request)
	return ret0
}

// GetRequest indicates an expected call of GetRequest
func (mr *MockResponseMockRecorder) GetRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockResponse)(nil).GetRequest))
}

// GetBody mocks base method
func (m *MockResponse) GetBody() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBody")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetBody indicates an expected call of GetBody
func (mr *MockResponseMockRecorder) GetBody() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBody", reflect.TypeOf((*MockResponse)(nil).GetBody))
}

// Abandon mocks base method
func (m *MockResponse) Abandon() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Abandon")
}

// Abandon indicates an expected call of Abandon
func (mr *MockResponseMockRecorder) Abandon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abandon", reflect.TypeOf((*MockResponse)(nil).Abandon))
}