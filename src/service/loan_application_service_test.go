package service

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/rxrav/loan_app/src/domain"
	"github.com/rxrav/loan_app/src/dto"
	mock_repo "github.com/rxrav/loan_app/src/mocks/repo"
	mock_service "github.com/rxrav/loan_app/src/mocks/service"
)

func TestApply_ShouldPanic(t *testing.T) {
	// arrange
	request := dto.LoanApplicationRequest{
		Username:      "maxpayne",
		AppliedAmount: float64(100_005),
	}
	mockCtrl := gomock.NewController(t)

	mockUserService := mock_service.NewMockUserService(mockCtrl)
	mockLoanApprovalService := mock_service.NewMockLoanApprovalService(mockCtrl)
	mockLoanRepo := mock_repo.NewMockLoanRepo(mockCtrl)

	// defer assert
	defer func() {
		want := "loan application amount is too high"
		if err := recover(); err != nil {
			got := err.(dto.LoanApplicationError)
			if got.ErrDetails != want {
				t.Errorf("got: %s, want: %s", got.ErrDetails, want)
			}
		}
	}()

	// act
	loanApplicationService := newLoanApplicationService(mockUserService, mockLoanApprovalService, mockLoanRepo)
	_ = loanApplicationService.Apply(request)

	// assert
	t.Errorf("should have panicked")
}

func TestApply(t *testing.T) {
	// arrange
	want := (100_000 * 0.7) * 0.8
	request := dto.LoanApplicationRequest{
		Username:      "maxpayne",
		AppliedAmount: float64(100_000),
	}

	mockCtrl := gomock.NewController(t)

	mockUserService := mock_service.NewMockUserService(mockCtrl)
	mockUserService.EXPECT().GetUser("maxpayne").Return(&dto.UserDetails{
		UserID:       uuid.New().String(),
		Username:     "maxpayne987654",
		FirstName:    "Max",
		LastName:     "Payne",
		Age:          42,
		SocialNumber: "MAX123456",
	}).Times(1)

	mockLoanApprovalService := mock_service.NewMockLoanApprovalService(mockCtrl)
	mockLoanApprovalService.EXPECT().Validate("MAX123456", float64(100_000)).Return(100_000 * 0.7).Times(1)
	mockLoanApprovalService.EXPECT().Sanction(42, float64(70_000)).Return(true, (100_000*0.7)*0.8).Times(1)

	mockLoanRepo := mock_repo.NewMockLoanRepo(mockCtrl)
	mockLoanRepo.EXPECT().CreateLoan(domain.Loan{
		Username:       "maxpayne",
		AppliedAmount:  float64(100_000),
		ApprovedAmount: (100_000 * 0.7) * 0.8,
		Status:         "Approved",
	}).Times(1)

	// act
	loanApplicationService := newLoanApplicationService(mockUserService, mockLoanApprovalService, mockLoanRepo)
	got := loanApplicationService.Apply(request)

	// assert
	if got.ApprovedAmount != want {
		_ = fmt.Errorf("wanted %.2f, got %.2f", want, got.AppliedAmount)
	}
}

func TestGetLoan(t *testing.T) {
	// Coding Exercise
}

func TestGetAllLoans(t *testing.T) {
	// Coding Exercise
}
