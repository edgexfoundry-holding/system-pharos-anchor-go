// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCommand is a mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// AddApp mocks base method
func (m *MockCommand) AddApp(appId string, description []byte) error {
	ret := m.ctrl.Call(m, "AddApp", appId, description)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddApp indicates an expected call of AddApp
func (mr *MockCommandMockRecorder) AddApp(appId, description interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddApp", reflect.TypeOf((*MockCommand)(nil).AddApp), appId, description)
}

// GetApp mocks base method
func (m *MockCommand) GetApp(appId string) (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "GetApp", appId)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApp indicates an expected call of GetApp
func (mr *MockCommandMockRecorder) GetApp(appId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockCommand)(nil).GetApp), appId)
}

// GetApps mocks base method
func (m *MockCommand) GetApps(queryOptional ...map[string]interface{}) ([]map[string]interface{}, error) {
	varargs := []interface{}{}
	for _, a := range queryOptional {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApps", varargs...)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApps indicates an expected call of GetApps
func (mr *MockCommandMockRecorder) GetApps(queryOptional ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApps", reflect.TypeOf((*MockCommand)(nil).GetApps), queryOptional...)
}

// DeleteApp mocks base method
func (m *MockCommand) DeleteApp(appId string) error {
	ret := m.ctrl.Call(m, "DeleteApp", appId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApp indicates an expected call of DeleteApp
func (mr *MockCommandMockRecorder) DeleteApp(appId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApp", reflect.TypeOf((*MockCommand)(nil).DeleteApp), appId)
}
