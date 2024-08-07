// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/manifests/api (interfaces: ClusterManifestsInternals)

// Package api is a generated GoMock package.
package api

import (
	context "context"
	reflect "reflect"

	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
	models "github.com/openshift/assisted-service/models"
	manifests "github.com/openshift/assisted-service/restapi/operations/manifests"
)

// MockClusterManifestsInternals is a mock of ClusterManifestsInternals interface.
type MockClusterManifestsInternals struct {
	ctrl     *gomock.Controller
	recorder *MockClusterManifestsInternalsMockRecorder
}

// MockClusterManifestsInternalsMockRecorder is the mock recorder for MockClusterManifestsInternals.
type MockClusterManifestsInternalsMockRecorder struct {
	mock *MockClusterManifestsInternals
}

// NewMockClusterManifestsInternals creates a new mock instance.
func NewMockClusterManifestsInternals(ctrl *gomock.Controller) *MockClusterManifestsInternals {
	mock := &MockClusterManifestsInternals{ctrl: ctrl}
	mock.recorder = &MockClusterManifestsInternalsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterManifestsInternals) EXPECT() *MockClusterManifestsInternalsMockRecorder {
	return m.recorder
}

// CreateClusterManifestInternal mocks base method.
func (m *MockClusterManifestsInternals) CreateClusterManifestInternal(arg0 context.Context, arg1 manifests.V2CreateClusterManifestParams, arg2 bool) (*models.Manifest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClusterManifestInternal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClusterManifestInternal indicates an expected call of CreateClusterManifestInternal.
func (mr *MockClusterManifestsInternalsMockRecorder) CreateClusterManifestInternal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClusterManifestInternal", reflect.TypeOf((*MockClusterManifestsInternals)(nil).CreateClusterManifestInternal), arg0, arg1, arg2)
}

// DeleteClusterManifestInternal mocks base method.
func (m *MockClusterManifestsInternals) DeleteClusterManifestInternal(arg0 context.Context, arg1 manifests.V2DeleteClusterManifestParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClusterManifestInternal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClusterManifestInternal indicates an expected call of DeleteClusterManifestInternal.
func (mr *MockClusterManifestsInternalsMockRecorder) DeleteClusterManifestInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClusterManifestInternal", reflect.TypeOf((*MockClusterManifestsInternals)(nil).DeleteClusterManifestInternal), arg0, arg1)
}

// FindUserManifestPathsByLegacyMetadata mocks base method.
func (m *MockClusterManifestsInternals) FindUserManifestPathsByLegacyMetadata(arg0 context.Context, arg1 strfmt.UUID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserManifestPathsByLegacyMetadata", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserManifestPathsByLegacyMetadata indicates an expected call of FindUserManifestPathsByLegacyMetadata.
func (mr *MockClusterManifestsInternalsMockRecorder) FindUserManifestPathsByLegacyMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserManifestPathsByLegacyMetadata", reflect.TypeOf((*MockClusterManifestsInternals)(nil).FindUserManifestPathsByLegacyMetadata), arg0, arg1)
}

// ListClusterManifestsInternal mocks base method.
func (m *MockClusterManifestsInternals) ListClusterManifestsInternal(arg0 context.Context, arg1 manifests.V2ListClusterManifestsParams) (models.ListManifests, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusterManifestsInternal", arg0, arg1)
	ret0, _ := ret[0].(models.ListManifests)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusterManifestsInternal indicates an expected call of ListClusterManifestsInternal.
func (mr *MockClusterManifestsInternalsMockRecorder) ListClusterManifestsInternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusterManifestsInternal", reflect.TypeOf((*MockClusterManifestsInternals)(nil).ListClusterManifestsInternal), arg0, arg1)
}
