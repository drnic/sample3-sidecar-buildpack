// Code generated by MockGen. DO NOT EDIT.
// Source: hooks.go

// Package libbuildpack_test is a generated GoMock package.
package libbuildpack_test

import (
	. "github.com/cloudfoundry/libbuildpack"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHook is a mock of Hook interface
type MockHook struct {
	ctrl     *gomock.Controller
	recorder *MockHookMockRecorder
}

// MockHookMockRecorder is the mock recorder for MockHook
type MockHookMockRecorder struct {
	mock *MockHook
}

// NewMockHook creates a new mock instance
func NewMockHook(ctrl *gomock.Controller) *MockHook {
	mock := &MockHook{ctrl: ctrl}
	mock.recorder = &MockHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHook) EXPECT() *MockHookMockRecorder {
	return m.recorder
}

// BeforeCompile mocks base method
func (m *MockHook) BeforeCompile(arg0 *Stager) error {
	ret := m.ctrl.Call(m, "BeforeCompile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeCompile indicates an expected call of BeforeCompile
func (mr *MockHookMockRecorder) BeforeCompile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeCompile", reflect.TypeOf((*MockHook)(nil).BeforeCompile), arg0)
}

// AfterCompile mocks base method
func (m *MockHook) AfterCompile(arg0 *Stager) error {
	ret := m.ctrl.Call(m, "AfterCompile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterCompile indicates an expected call of AfterCompile
func (mr *MockHookMockRecorder) AfterCompile(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterCompile", reflect.TypeOf((*MockHook)(nil).AfterCompile), arg0)
}
