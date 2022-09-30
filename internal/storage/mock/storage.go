// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	storage "github.com/Decentr-net/vulcan/internal/storage"
	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// InTx mocks base method
func (m *MockStorage) InTx(ctx context.Context, f func(storage.Storage) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InTx", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// InTx indicates an expected call of InTx
func (mr *MockStorageMockRecorder) InTx(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InTx", reflect.TypeOf((*MockStorage)(nil).InTx), ctx, f)
}

// GetConfirmedRegistrationsTotal mocks base method
func (m *MockStorage) GetConfirmedRegistrationsTotal(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfirmedRegistrationsTotal", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfirmedRegistrationsTotal indicates an expected call of GetConfirmedRegistrationsTotal
func (mr *MockStorageMockRecorder) GetConfirmedRegistrationsTotal(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfirmedRegistrationsTotal", reflect.TypeOf((*MockStorage)(nil).GetConfirmedRegistrationsTotal), ctx)
}

// GetConfirmedRegistrationsStats mocks base method
func (m *MockStorage) GetConfirmedRegistrationsStats(ctx context.Context) ([]*storage.RegisterStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfirmedRegistrationsStats", ctx)
	ret0, _ := ret[0].([]*storage.RegisterStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfirmedRegistrationsStats indicates an expected call of GetConfirmedRegistrationsStats
func (mr *MockStorageMockRecorder) GetConfirmedRegistrationsStats(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfirmedRegistrationsStats", reflect.TypeOf((*MockStorage)(nil).GetConfirmedRegistrationsStats), ctx)
}

// GetRequestByOwner mocks base method
func (m *MockStorage) GetRequestByOwner(ctx context.Context, owner string) (*storage.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestByOwner", ctx, owner)
	ret0, _ := ret[0].(*storage.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestByOwner indicates an expected call of GetRequestByOwner
func (mr *MockStorageMockRecorder) GetRequestByOwner(ctx, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestByOwner", reflect.TypeOf((*MockStorage)(nil).GetRequestByOwner), ctx, owner)
}

// GetRequestByOwnReferralCode mocks base method
func (m *MockStorage) GetRequestByOwnReferralCode(ctx context.Context, ownReferralCode string) (*storage.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestByOwnReferralCode", ctx, ownReferralCode)
	ret0, _ := ret[0].(*storage.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestByOwnReferralCode indicates an expected call of GetRequestByOwnReferralCode
func (mr *MockStorageMockRecorder) GetRequestByOwnReferralCode(ctx, ownReferralCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestByOwnReferralCode", reflect.TypeOf((*MockStorage)(nil).GetRequestByOwnReferralCode), ctx, ownReferralCode)
}

// GetRequestByAddress mocks base method
func (m *MockStorage) GetRequestByAddress(ctx context.Context, address string) (*storage.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestByAddress", ctx, address)
	ret0, _ := ret[0].(*storage.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestByAddress indicates an expected call of GetRequestByAddress
func (mr *MockStorageMockRecorder) GetRequestByAddress(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestByAddress", reflect.TypeOf((*MockStorage)(nil).GetRequestByAddress), ctx, address)
}

// SetConfirmed mocks base method
func (m *MockStorage) SetConfirmed(ctx context.Context, owner string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfirmed", ctx, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfirmed indicates an expected call of SetConfirmed
func (mr *MockStorageMockRecorder) SetConfirmed(ctx, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfirmed", reflect.TypeOf((*MockStorage)(nil).SetConfirmed), ctx, owner)
}

// CreateTestnetConfirmedRequest mocks base method
func (m *MockStorage) CreateTestnetConfirmedRequest(ctx context.Context, address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestnetConfirmedRequest", ctx, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTestnetConfirmedRequest indicates an expected call of CreateTestnetConfirmedRequest
func (mr *MockStorageMockRecorder) CreateTestnetConfirmedRequest(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestnetConfirmedRequest", reflect.TypeOf((*MockStorage)(nil).CreateTestnetConfirmedRequest), ctx, address)
}

// UpsertRequest mocks base method
func (m *MockStorage) UpsertRequest(ctx context.Context, owner, email, address, code string, referralCode sql.NullString) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRequest", ctx, owner, email, address, code, referralCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRequest indicates an expected call of UpsertRequest
func (mr *MockStorageMockRecorder) UpsertRequest(ctx, owner, email, address, code, referralCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRequest", reflect.TypeOf((*MockStorage)(nil).UpsertRequest), ctx, owner, email, address, code, referralCode)
}

// CreateReferralTracking mocks base method
func (m *MockStorage) CreateReferralTracking(ctx context.Context, receiver, referralCode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReferralTracking", ctx, receiver, referralCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReferralTracking indicates an expected call of CreateReferralTracking
func (mr *MockStorageMockRecorder) CreateReferralTracking(ctx, receiver, referralCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReferralTracking", reflect.TypeOf((*MockStorage)(nil).CreateReferralTracking), ctx, receiver, referralCode)
}

// TransitionReferralTrackingToInstalled mocks base method
func (m *MockStorage) TransitionReferralTrackingToInstalled(ctx context.Context, receiver string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransitionReferralTrackingToInstalled", ctx, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransitionReferralTrackingToInstalled indicates an expected call of TransitionReferralTrackingToInstalled
func (mr *MockStorageMockRecorder) TransitionReferralTrackingToInstalled(ctx, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransitionReferralTrackingToInstalled", reflect.TypeOf((*MockStorage)(nil).TransitionReferralTrackingToInstalled), ctx, receiver)
}

// TransitionReferralTrackingToConfirmed mocks base method
func (m *MockStorage) TransitionReferralTrackingToConfirmed(ctx context.Context, receiver string, senderReward, receiverReward types.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransitionReferralTrackingToConfirmed", ctx, receiver, senderReward, receiverReward)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransitionReferralTrackingToConfirmed indicates an expected call of TransitionReferralTrackingToConfirmed
func (mr *MockStorageMockRecorder) TransitionReferralTrackingToConfirmed(ctx, receiver, senderReward, receiverReward interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransitionReferralTrackingToConfirmed", reflect.TypeOf((*MockStorage)(nil).TransitionReferralTrackingToConfirmed), ctx, receiver, senderReward, receiverReward)
}

// GetReferralTrackingByReceiver mocks base method
func (m *MockStorage) GetReferralTrackingByReceiver(ctx context.Context, receiver string) (*storage.ReferralTracking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReferralTrackingByReceiver", ctx, receiver)
	ret0, _ := ret[0].(*storage.ReferralTracking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReferralTrackingByReceiver indicates an expected call of GetReferralTrackingByReceiver
func (mr *MockStorageMockRecorder) GetReferralTrackingByReceiver(ctx, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReferralTrackingByReceiver", reflect.TypeOf((*MockStorage)(nil).GetReferralTrackingByReceiver), ctx, receiver)
}

// GetReferralTrackingStats mocks base method
func (m *MockStorage) GetReferralTrackingStats(ctx context.Context, sender string) ([]*storage.ReferralTrackingStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReferralTrackingStats", ctx, sender)
	ret0, _ := ret[0].([]*storage.ReferralTrackingStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReferralTrackingStats indicates an expected call of GetReferralTrackingStats
func (mr *MockStorageMockRecorder) GetReferralTrackingStats(ctx, sender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReferralTrackingStats", reflect.TypeOf((*MockStorage)(nil).GetReferralTrackingStats), ctx, sender)
}

// GetUnconfirmedReferralTracking mocks base method
func (m *MockStorage) GetUnconfirmedReferralTracking(ctx context.Context, days int) ([]*storage.ReferralTracking, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnconfirmedReferralTracking", ctx, days)
	ret0, _ := ret[0].([]*storage.ReferralTracking)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnconfirmedReferralTracking indicates an expected call of GetUnconfirmedReferralTracking
func (mr *MockStorageMockRecorder) GetUnconfirmedReferralTracking(ctx, days interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnconfirmedReferralTracking", reflect.TypeOf((*MockStorage)(nil).GetUnconfirmedReferralTracking), ctx, days)
}

// GetConfirmedReferralTrackingCount mocks base method
func (m *MockStorage) GetConfirmedReferralTrackingCount(ctx context.Context, sender string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfirmedReferralTrackingCount", ctx, sender)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfirmedReferralTrackingCount indicates an expected call of GetConfirmedReferralTrackingCount
func (mr *MockStorageMockRecorder) GetConfirmedReferralTrackingCount(ctx, sender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfirmedReferralTrackingCount", reflect.TypeOf((*MockStorage)(nil).GetConfirmedReferralTrackingCount), ctx, sender)
}

// DoesEmailHaveFraudDomain mocks base method
func (m *MockStorage) DoesEmailHaveFraudDomain(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesEmailHaveFraudDomain", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesEmailHaveFraudDomain indicates an expected call of DoesEmailHaveFraudDomain
func (mr *MockStorageMockRecorder) DoesEmailHaveFraudDomain(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesEmailHaveFraudDomain", reflect.TypeOf((*MockStorage)(nil).DoesEmailHaveFraudDomain), ctx, email)
}

// CreateDLoan mocks base method
func (m *MockStorage) CreateDLoan(ctx context.Context, address, firstName, lastName string, pdv float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDLoan", ctx, address, firstName, lastName, pdv)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDLoan indicates an expected call of CreateDLoan
func (mr *MockStorageMockRecorder) CreateDLoan(ctx, address, firstName, lastName, pdv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDLoan", reflect.TypeOf((*MockStorage)(nil).CreateDLoan), ctx, address, firstName, lastName, pdv)
}

// GetDLoans mocks base method
func (m *MockStorage) GetDLoans(ctx context.Context, take, skip int) ([]*storage.DLoan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDLoans", ctx, take, skip)
	ret0, _ := ret[0].([]*storage.DLoan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDLoans indicates an expected call of GetDLoans
func (mr *MockStorageMockRecorder) GetDLoans(ctx, take, skip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDLoans", reflect.TypeOf((*MockStorage)(nil).GetDLoans), ctx, take, skip)
}
