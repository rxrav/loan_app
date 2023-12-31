// Code generated by MockGen. DO NOT EDIT.
// Source: .\src\service\loan_approval_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLoanApprovalService is a mock of LoanApprovalService interface.
type MockLoanApprovalService struct {
	ctrl     *gomock.Controller
	recorder *MockLoanApprovalServiceMockRecorder
}

// MockLoanApprovalServiceMockRecorder is the mock recorder for MockLoanApprovalService.
type MockLoanApprovalServiceMockRecorder struct {
	mock *MockLoanApprovalService
}

// NewMockLoanApprovalService creates a new mock instance.
func NewMockLoanApprovalService(ctrl *gomock.Controller) *MockLoanApprovalService {
	mock := &MockLoanApprovalService{ctrl: ctrl}
	mock.recorder = &MockLoanApprovalServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoanApprovalService) EXPECT() *MockLoanApprovalServiceMockRecorder {
	return m.recorder
}

// Sanction mocks base method.
func (m *MockLoanApprovalService) Sanction(age int, maxPossibleAmount float64) (bool, float64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sanction", age, maxPossibleAmount)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(float64)
	return ret0, ret1
}

// Sanction indicates an expected call of Sanction.
func (mr *MockLoanApprovalServiceMockRecorder) Sanction(age, maxPossibleAmount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sanction", reflect.TypeOf((*MockLoanApprovalService)(nil).Sanction), age, maxPossibleAmount)
}

// Validate mocks base method.
func (m *MockLoanApprovalService) Validate(socialNumber string, appliedAmount float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", socialNumber, appliedAmount)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockLoanApprovalServiceMockRecorder) Validate(socialNumber, appliedAmount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockLoanApprovalService)(nil).Validate), socialNumber, appliedAmount)
}