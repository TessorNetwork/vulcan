// Code generated by MockGen. DO NOT EDIT.
// Source: sender.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSender is a mock of Sender interface
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// SendVerificationEmailAsync mocks base method
func (m *MockSender) SendVerificationEmailAsync(ctx context.Context, email, code string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendVerificationEmailAsync", ctx, email, code)
}

// SendVerificationEmailAsync indicates an expected call of SendVerificationEmailAsync
func (mr *MockSenderMockRecorder) SendVerificationEmailAsync(ctx, email, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerificationEmailAsync", reflect.TypeOf((*MockSender)(nil).SendVerificationEmailAsync), ctx, email, code)
}

// SendWelcomeEmailAsync mocks base method
func (m *MockSender) SendWelcomeEmailAsync(ctx context.Context, email string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendWelcomeEmailAsync", ctx, email)
}

// SendWelcomeEmailAsync indicates an expected call of SendWelcomeEmailAsync
func (mr *MockSenderMockRecorder) SendWelcomeEmailAsync(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendWelcomeEmailAsync", reflect.TypeOf((*MockSender)(nil).SendWelcomeEmailAsync), ctx, email)
}
