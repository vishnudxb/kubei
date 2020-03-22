// Code generated by MockGen. DO NOT EDIT.
// Source: kubei/common (interfaces: K8ContextSecretServiceInterface)

// Package common is a generated GoMock package.
package common

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
)

// MockK8ContextSecretServiceInterface is a mock of K8ContextSecretServiceInterface interface
type MockK8ContextSecretServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockK8ContextSecretServiceInterfaceMockRecorder
}

// MockK8ContextSecretServiceInterfaceMockRecorder is the mock recorder for MockK8ContextSecretServiceInterface
type MockK8ContextSecretServiceInterfaceMockRecorder struct {
	mock *MockK8ContextSecretServiceInterface
}

// NewMockK8ContextSecretServiceInterface creates a new mock instance
func NewMockK8ContextSecretServiceInterface(ctrl *gomock.Controller) *MockK8ContextSecretServiceInterface {
	mock := &MockK8ContextSecretServiceInterface{ctrl: ctrl}
	mock.recorder = &MockK8ContextSecretServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockK8ContextSecretServiceInterface) EXPECT() *MockK8ContextSecretServiceInterfaceMockRecorder {
	return m.recorder
}

// GetMatchingSecretName mocks base method
func (m *MockK8ContextSecretServiceInterface) GetMatchingSecretName(arg0 []v1.Secret, arg1 v1.Container) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchingSecretName", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMatchingSecretName indicates an expected call of GetMatchingSecretName
func (mr *MockK8ContextSecretServiceInterfaceMockRecorder) GetMatchingSecretName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchingSecretName", reflect.TypeOf((*MockK8ContextSecretServiceInterface)(nil).GetMatchingSecretName), arg0, arg1)
}