/*******************************************************************************
 * Copyright 2017 Samsung Electronics All Rights Reserved.
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
// Source: groupcontroller.go

// Package mock_group is a generated GoMock package.
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

// DeployApp mocks base method
func (m *MockCommand) DeployApp(groupId, body string) (int, map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "DeployApp", groupId, body)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeployApp indicates an expected call of DeployApp
func (mr *MockCommandMockRecorder) DeployApp(groupId, body interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployApp", reflect.TypeOf((*MockCommand)(nil).DeployApp), groupId, body)
}

// GetApps mocks base method
func (m *MockCommand) GetApps(groupId string) (int, map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "GetApps", groupId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetApps indicates an expected call of GetApps
func (mr *MockCommandMockRecorder) GetApps(groupId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApps", reflect.TypeOf((*MockCommand)(nil).GetApps), groupId)
}

// GetApp mocks base method
func (m *MockCommand) GetApp(groupId, appId string) (int, map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "GetApp", groupId, appId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetApp indicates an expected call of GetApp
func (mr *MockCommandMockRecorder) GetApp(groupId, appId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockCommand)(nil).GetApp), groupId, appId)
}

// UpdateAppInfo mocks base method
func (m *MockCommand) UpdateAppInfo(groupId, appId, body string) (int, map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "UpdateAppInfo", groupId, appId, body)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateAppInfo indicates an expected call of UpdateAppInfo
func (mr *MockCommandMockRecorder) UpdateAppInfo(groupId, appId, body interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppInfo", reflect.TypeOf((*MockCommand)(nil).UpdateAppInfo), groupId, appId, body)
}

// DeleteApp mocks base method
func (m *MockCommand) DeleteApp(groupId, appId string) (int, map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "DeleteApp", groupId, appId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteApp indicates an expected call of DeleteApp
func (mr *MockCommandMockRecorder) DeleteApp(groupId, appId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApp", reflect.TypeOf((*MockCommand)(nil).DeleteApp), groupId, appId)
}

// UpdateApp mocks base method
func (m *MockCommand) UpdateApp(groupId, appId string) (int, map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "UpdateApp", groupId, appId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateApp indicates an expected call of UpdateApp
func (mr *MockCommandMockRecorder) UpdateApp(groupId, appId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApp", reflect.TypeOf((*MockCommand)(nil).UpdateApp), groupId, appId)
}

// StartApp mocks base method
func (m *MockCommand) StartApp(groupId, appId string) (int, map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "StartApp", groupId, appId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StartApp indicates an expected call of StartApp
func (mr *MockCommandMockRecorder) StartApp(groupId, appId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartApp", reflect.TypeOf((*MockCommand)(nil).StartApp), groupId, appId)
}

// StopApp mocks base method
func (m *MockCommand) StopApp(groupId, appId string) (int, map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "StopApp", groupId, appId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StopApp indicates an expected call of StopApp
func (mr *MockCommandMockRecorder) StopApp(groupId, appId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopApp", reflect.TypeOf((*MockCommand)(nil).StopApp), groupId, appId)
}
