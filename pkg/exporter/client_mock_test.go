// Code generated by MockGen. DO NOT EDIT.
// Source: ../client/client.go

// Package exporter is a generated GoMock package.
package exporter

import (
	reflect "reflect"

	"github.com/jamesnetherton/homehub-metrics-exporter/pkg/client"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockClient) Login() *client.Response {
	ret := m.ctrl.Call(m, "Login")
	ret0, _ := ret[0].(*client.Response)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockClientMockRecorder) Login() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockClient)(nil).Login))
}

// GetSummaryStatistics mocks base method
func (m *MockClient) GetSummaryStatistics() *client.Response {
	ret := m.ctrl.Call(m, "GetSummaryStatistics")
	ret0, _ := ret[0].(*client.Response)
	return ret0
}

// GetSummaryStatistics indicates an expected call of GetSummaryStatistics
func (mr *MockClientMockRecorder) GetSummaryStatistics() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSummaryStatistics", reflect.TypeOf((*MockClient)(nil).GetSummaryStatistics))
}

// GetBandwithStatistics mocks base method
func (m *MockClient) GetBandwithStatistics() *client.Response {
	ret := m.ctrl.Call(m, "GetBandwithStatistics")
	ret0, _ := ret[0].(*client.Response)
	return ret0
}

// GetBandwithStatistics indicates an expected call of GetBandwithStatistics
func (mr *MockClientMockRecorder) GetBandwithStatistics() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBandwithStatistics", reflect.TypeOf((*MockClient)(nil).GetBandwithStatistics))
}
