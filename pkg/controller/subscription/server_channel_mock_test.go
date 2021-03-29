// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/protocols/e2ap101/channels/ric_channel.go

// Package subscription is a generated GoMock package.
package subscription

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	e2ap_pdu_contents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
)

// MockRICChannel is a mock of RICChannel interface.
type MockRICChannel struct {
	ctrl     *gomock.Controller
	recorder *MockRICChannelMockRecorder
}

// MockRICChannelMockRecorder is the mock recorder for MockRICChannel.
type MockRICChannelMockRecorder struct {
	mock *MockRICChannel
}

// NewMockRICChannel creates a new mock instance.
func NewMockRICChannel(ctrl *gomock.Controller) *MockRICChannel {
	mock := &MockRICChannel{ctrl: ctrl}
	mock.recorder = &MockRICChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRICChannel) EXPECT() *MockRICChannelMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRICChannel) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRICChannelMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRICChannel)(nil).Close))
}

// Context mocks base method.
func (m *MockRICChannel) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockRICChannelMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockRICChannel)(nil).Context))
}

// LocalAddr mocks base method.
func (m *MockRICChannel) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockRICChannelMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockRICChannel)(nil).LocalAddr))
}

// RICControl mocks base method.
func (m *MockRICChannel) RICControl(ctx context.Context, request *e2ap_pdu_contents.RiccontrolRequest) (*e2ap_pdu_contents.RiccontrolAcknowledge, *e2ap_pdu_contents.RiccontrolFailure, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RICControl", ctx, request)
	ret0, _ := ret[0].(*e2ap_pdu_contents.RiccontrolAcknowledge)
	ret1, _ := ret[1].(*e2ap_pdu_contents.RiccontrolFailure)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RICControl indicates an expected call of RICControl.
func (mr *MockRICChannelMockRecorder) RICControl(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RICControl", reflect.TypeOf((*MockRICChannel)(nil).RICControl), ctx, request)
}

// RICSubscription mocks base method.
func (m *MockRICChannel) RICSubscription(ctx context.Context, request *e2ap_pdu_contents.RicsubscriptionRequest) (*e2ap_pdu_contents.RicsubscriptionResponse, *e2ap_pdu_contents.RicsubscriptionFailure, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RICSubscription", ctx, request)
	ret0, _ := ret[0].(*e2ap_pdu_contents.RicsubscriptionResponse)
	ret1, _ := ret[1].(*e2ap_pdu_contents.RicsubscriptionFailure)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RICSubscription indicates an expected call of RICSubscription.
func (mr *MockRICChannelMockRecorder) RICSubscription(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RICSubscription", reflect.TypeOf((*MockRICChannel)(nil).RICSubscription), ctx, request)
}

// RICSubscriptionDelete mocks base method.
func (m *MockRICChannel) RICSubscriptionDelete(ctx context.Context, request *e2ap_pdu_contents.RicsubscriptionDeleteRequest) (*e2ap_pdu_contents.RicsubscriptionDeleteResponse, *e2ap_pdu_contents.RicsubscriptionDeleteFailure, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RICSubscriptionDelete", ctx, request)
	ret0, _ := ret[0].(*e2ap_pdu_contents.RicsubscriptionDeleteResponse)
	ret1, _ := ret[1].(*e2ap_pdu_contents.RicsubscriptionDeleteFailure)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RICSubscriptionDelete indicates an expected call of RICSubscriptionDelete.
func (mr *MockRICChannelMockRecorder) RICSubscriptionDelete(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RICSubscriptionDelete", reflect.TypeOf((*MockRICChannel)(nil).RICSubscriptionDelete), ctx, request)
}

// RemoteAddr mocks base method.
func (m *MockRICChannel) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockRICChannelMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockRICChannel)(nil).RemoteAddr))
}
