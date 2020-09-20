// Code generated by MockGen. DO NOT EDIT.
// Source: blockchain.go

// Package blockchain is a generated GoMock package.
package blockchain

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBlockchain is a mock of Blockchain interface
type MockBlockchain struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainMockRecorder
}

// MockBlockchainMockRecorder is the mock recorder for MockBlockchain
type MockBlockchainMockRecorder struct {
	mock *MockBlockchain
}

// NewMockBlockchain creates a new mock instance
func NewMockBlockchain(ctrl *gomock.Controller) *MockBlockchain {
	mock := &MockBlockchain{ctrl: ctrl}
	mock.recorder = &MockBlockchainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockchain) EXPECT() *MockBlockchainMockRecorder {
	return m.recorder
}

// SendStakes mocks base method
func (m *MockBlockchain) SendStakes(address string, amount int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendStakes", address, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendStakes indicates an expected call of SendStakes
func (mr *MockBlockchainMockRecorder) SendStakes(address, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendStakes", reflect.TypeOf((*MockBlockchain)(nil).SendStakes), address, amount)
}
