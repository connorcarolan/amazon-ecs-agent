// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/utils/ttime (interfaces: Time,Timer)

// Package mock_ttime is a generated GoMock package.
package mock_ttime

import (
	reflect "reflect"
	time "time"

	ttime "github.com/aws/amazon-ecs-agent/agent/utils/ttime"
	gomock "github.com/golang/mock/gomock"
)

// MockTime is a mock of Time interface
type MockTime struct {
	ctrl     *gomock.Controller
	recorder *MockTimeMockRecorder
}

// MockTimeMockRecorder is the mock recorder for MockTime
type MockTimeMockRecorder struct {
	mock *MockTime
}

// NewMockTime creates a new mock instance
func NewMockTime(ctrl *gomock.Controller) *MockTime {
	mock := &MockTime{ctrl: ctrl}
	mock.recorder = &MockTimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTime) EXPECT() *MockTimeMockRecorder {
	return m.recorder
}

// After mocks base method
func (m *MockTime) After(arg0 time.Duration) <-chan time.Time {
	ret := m.ctrl.Call(m, "After", arg0)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// After indicates an expected call of After
func (mr *MockTimeMockRecorder) After(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "After", reflect.TypeOf((*MockTime)(nil).After), arg0)
}

// AfterFunc mocks base method
func (m *MockTime) AfterFunc(arg0 time.Duration, arg1 func()) ttime.Timer {
	ret := m.ctrl.Call(m, "AfterFunc", arg0, arg1)
	ret0, _ := ret[0].(ttime.Timer)
	return ret0
}

// AfterFunc indicates an expected call of AfterFunc
func (mr *MockTimeMockRecorder) AfterFunc(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterFunc", reflect.TypeOf((*MockTime)(nil).AfterFunc), arg0, arg1)
}

// Now mocks base method
func (m *MockTime) Now() time.Time {
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now
func (mr *MockTimeMockRecorder) Now() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockTime)(nil).Now))
}

// Sleep mocks base method
func (m *MockTime) Sleep(arg0 time.Duration) {
	m.ctrl.Call(m, "Sleep", arg0)
}

// Sleep indicates an expected call of Sleep
func (mr *MockTimeMockRecorder) Sleep(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MockTime)(nil).Sleep), arg0)
}

// MockTimer is a mock of Timer interface
type MockTimer struct {
	ctrl     *gomock.Controller
	recorder *MockTimerMockRecorder
}

// MockTimerMockRecorder is the mock recorder for MockTimer
type MockTimerMockRecorder struct {
	mock *MockTimer
}

// NewMockTimer creates a new mock instance
func NewMockTimer(ctrl *gomock.Controller) *MockTimer {
	mock := &MockTimer{ctrl: ctrl}
	mock.recorder = &MockTimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimer) EXPECT() *MockTimerMockRecorder {
	return m.recorder
}

// Reset mocks base method
func (m *MockTimer) Reset(arg0 time.Duration) bool {
	ret := m.ctrl.Call(m, "Reset", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Reset indicates an expected call of Reset
func (mr *MockTimerMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockTimer)(nil).Reset), arg0)
}

// Stop mocks base method
func (m *MockTimer) Stop() bool {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockTimerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTimer)(nil).Stop))
}
