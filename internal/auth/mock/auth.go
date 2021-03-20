// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSecretStore is a mock of SecretStore interface
type MockSecretStore struct {
	ctrl     *gomock.Controller
	recorder *MockSecretStoreMockRecorder
}

// MockSecretStoreMockRecorder is the mock recorder for MockSecretStore
type MockSecretStoreMockRecorder struct {
	mock *MockSecretStore
}

// NewMockSecretStore creates a new mock instance
func NewMockSecretStore(ctrl *gomock.Controller) *MockSecretStore {
	mock := &MockSecretStore{ctrl: ctrl}
	mock.recorder = &MockSecretStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSecretStore) EXPECT() *MockSecretStoreMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockSecretStore) Set(namespace, key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", namespace, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockSecretStoreMockRecorder) Set(namespace, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSecretStore)(nil).Set), namespace, key, value)
}

// Get mocks base method
func (m *MockSecretStore) Get(namespace, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", namespace, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSecretStoreMockRecorder) Get(namespace, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSecretStore)(nil).Get), namespace, key)
}

// Delete mocks base method
func (m *MockSecretStore) Delete(namespace, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", namespace, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockSecretStoreMockRecorder) Delete(namespace, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSecretStore)(nil).Delete), namespace, key)
}