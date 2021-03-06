// Automatically generated by MockGen. DO NOT EDIT!
// Source: node.go

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

// SearchNodes mocks base method
func (m *MockCommand) SearchNodes(query map[string][]string) (int, map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "SearchNodes", query)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchNodes indicates an expected call of SearchNodes
func (mr *MockCommandMockRecorder) SearchNodes(query interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodes", reflect.TypeOf((*MockCommand)(nil).SearchNodes), query)
}
