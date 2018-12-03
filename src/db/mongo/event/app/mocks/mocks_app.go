/*******************************************************************************
 * Copyright 2018 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package mock_app is a generated GoMock package.
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

// AddEvent mocks base method
func (m *MockCommand) AddEvent(id, subscriberId string, nodeId []string) error {
	ret := m.ctrl.Call(m, "AddEvent", id, subscriberId, nodeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEvent indicates an expected call of AddEvent
func (mr *MockCommandMockRecorder) AddEvent(id, subscriberId, nodeId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEvent", reflect.TypeOf((*MockCommand)(nil).AddEvent), id, subscriberId, nodeId)
}

// GetEvent mocks base method
func (m *MockCommand) GetEvent(id string) (map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "GetEvent", id)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent
func (mr *MockCommandMockRecorder) GetEvent(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockCommand)(nil).GetEvent), id)
}

// DeleteEvent mocks base method
func (m *MockCommand) DeleteEvent(id string) error {
	ret := m.ctrl.Call(m, "DeleteEvent", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEvent indicates an expected call of DeleteEvent
func (mr *MockCommandMockRecorder) DeleteEvent(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockCommand)(nil).DeleteEvent), id)
}

// UnRegisterEvent mocks base method
func (m *MockCommand) UnRegisterEvent(id, subscriberId string) error {
	ret := m.ctrl.Call(m, "UnRegisterEvent", id, subscriberId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnRegisterEvent indicates an expected call of UnRegisterEvent
func (mr *MockCommandMockRecorder) UnRegisterEvent(id, subscriberId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnRegisterEvent", reflect.TypeOf((*MockCommand)(nil).UnRegisterEvent), id, subscriberId)
}