// Code generated by MockGen. DO NOT EDIT.
// Source: uploader.go

// Package uploader is a generated GoMock package.
package uploader

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	api "github.com/openshift/assisted-service/internal/events/api"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// IsEnabled mocks base method.
func (m *MockClient) IsEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEnabled indicates an expected call of IsEnabled.
func (mr *MockClientMockRecorder) IsEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnabled", reflect.TypeOf((*MockClient)(nil).IsEnabled))
}

// UploadEvents mocks base method.
func (m *MockClient) UploadEvents(ctx context.Context, cluster *common.Cluster, eventsHandler api.Handler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadEvents", ctx, cluster, eventsHandler)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadEvents indicates an expected call of UploadEvents.
func (mr *MockClientMockRecorder) UploadEvents(ctx, cluster, eventsHandler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadEvents", reflect.TypeOf((*MockClient)(nil).UploadEvents), ctx, cluster, eventsHandler)
}