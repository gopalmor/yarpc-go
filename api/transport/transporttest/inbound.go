// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/api/transport (interfaces: Inbound)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package transporttest is a generated GoMock package.
package transporttest

import (
	gomock "github.com/golang/mock/gomock"
	transport "go.uber.org/yarpc/api/transport"
	reflect "reflect"
)

// MockInbound is a mock of Inbound interface
type MockInbound struct {
	ctrl     *gomock.Controller
	recorder *MockInboundMockRecorder
}

// MockInboundMockRecorder is the mock recorder for MockInbound
type MockInboundMockRecorder struct {
	mock *MockInbound
}

// NewMockInbound creates a new mock instance
func NewMockInbound(ctrl *gomock.Controller) *MockInbound {
	mock := &MockInbound{ctrl: ctrl}
	mock.recorder = &MockInboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInbound) EXPECT() *MockInboundMockRecorder {
	return m.recorder
}

// IsRunning mocks base method
func (m *MockInbound) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning
func (mr *MockInboundMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockInbound)(nil).IsRunning))
}

// SetRouter mocks base method
func (m *MockInbound) SetRouter(arg0 transport.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRouter", arg0)
}

// SetRouter indicates an expected call of SetRouter
func (mr *MockInboundMockRecorder) SetRouter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRouter", reflect.TypeOf((*MockInbound)(nil).SetRouter), arg0)
}

// Start mocks base method
func (m *MockInbound) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockInboundMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockInbound)(nil).Start))
}

// Stop mocks base method
func (m *MockInbound) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockInboundMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockInbound)(nil).Stop))
}

// Transports mocks base method
func (m *MockInbound) Transports() []transport.Transport {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transports")
	ret0, _ := ret[0].([]transport.Transport)
	return ret0
}

// Transports indicates an expected call of Transports
func (mr *MockInboundMockRecorder) Transports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transports", reflect.TypeOf((*MockInbound)(nil).Transports))
}
