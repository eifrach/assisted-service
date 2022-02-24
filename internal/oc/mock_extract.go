// Code generated by MockGen. DO NOT EDIT.
// Source: extract.go

// Package oc is a generated GoMock package.
package oc

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
)

// MockExtracter is a mock of Extracter interface.
type MockExtracter struct {
	ctrl     *gomock.Controller
	recorder *MockExtracterMockRecorder
}

// MockExtracterMockRecorder is the mock recorder for MockExtracter.
type MockExtracterMockRecorder struct {
	mock *MockExtracter
}

// NewMockExtracter creates a new mock instance.
func NewMockExtracter(ctrl *gomock.Controller) *MockExtracter {
	mock := &MockExtracter{ctrl: ctrl}
	mock.recorder = &MockExtracterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtracter) EXPECT() *MockExtracterMockRecorder {
	return m.recorder
}

// Extract mocks base method.
func (m *MockExtracter) Extract(log logrus.FieldLogger, imageIndexPath, openshiftVersion, filePath, pullSecret string, insecure bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Extract", log, imageIndexPath, openshiftVersion, filePath, pullSecret, insecure)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Extract indicates an expected call of Extract.
func (mr *MockExtracterMockRecorder) Extract(log, imageIndexPath, openshiftVersion, filePath, pullSecret, insecure interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Extract", reflect.TypeOf((*MockExtracter)(nil).Extract), log, imageIndexPath, openshiftVersion, filePath, pullSecret, insecure)
}

// ExtractDatabaseIndex mocks base method.
func (m *MockExtracter) ExtractDatabaseIndex(log logrus.FieldLogger, releaseImageMirror, openshiftVersion, pullSecret string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractDatabaseIndex", log, releaseImageMirror, openshiftVersion, pullSecret)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtractDatabaseIndex indicates an expected call of ExtractDatabaseIndex.
func (mr *MockExtracterMockRecorder) ExtractDatabaseIndex(log, releaseImageMirror, openshiftVersion, pullSecret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractDatabaseIndex", reflect.TypeOf((*MockExtracter)(nil).ExtractDatabaseIndex), log, releaseImageMirror, openshiftVersion, pullSecret)
}