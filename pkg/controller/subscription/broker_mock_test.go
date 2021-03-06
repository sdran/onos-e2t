// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/broker/subscription/broker.go

// Package subscription is a generated GoMock package.
package subscription

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	subscription "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	subscription0 "github.com/onosproject/onos-e2t/pkg/broker/subscription"
)

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockBroker) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockBrokerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBroker)(nil).Close))
}

// CloseStream mocks base method.
func (m *MockBroker) CloseStream(id subscription.ID) (subscription0.StreamReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseStream", id)
	ret0, _ := ret[0].(subscription0.StreamReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseStream indicates an expected call of CloseStream.
func (mr *MockBrokerMockRecorder) CloseStream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseStream", reflect.TypeOf((*MockBroker)(nil).CloseStream), id)
}

// GetStream mocks base method.
func (m *MockBroker) GetStream(id subscription0.StreamID) (subscription0.StreamWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStream", id)
	ret0, _ := ret[0].(subscription0.StreamWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStream indicates an expected call of GetStream.
func (mr *MockBrokerMockRecorder) GetStream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStream", reflect.TypeOf((*MockBroker)(nil).GetStream), id)
}

// OpenStream mocks base method.
func (m *MockBroker) OpenStream(id subscription.ID) (subscription0.StreamReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream", id)
	ret0, _ := ret[0].(subscription0.StreamReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockBrokerMockRecorder) OpenStream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockBroker)(nil).OpenStream), id)
}
