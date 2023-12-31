package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_client "github.com/rxrav/loan_app/src/mocks/client"
)

func TestValidate(t *testing.T) {
	// arrange
	socialNumber := "XXX12345"
	appliedAmount := float64(100_000)
	mockCreditScore := 850
	want := float64(90_000)

	mockCtrl := gomock.NewController(t)
	mockCreditScoreClient := mock_client.NewMockCreditScoreClient(mockCtrl)
	mockCreditScoreClient.EXPECT().GetScore(socialNumber).Return(mockCreditScore, nil).Times(1)

	service := newLoanApprovalService(mockCreditScoreClient)

	// act
	got := service.Validate(socialNumber, appliedAmount)

	// assert
	if want != got {
		t.Errorf("wanted %.2f, got %.2f", want, got)
	}
}

func TestSanction(t *testing.T) {
	// arrange
	age := 35
	maxPossibleAmount := 100_000
	want := float64(80_000)

	mockCtrl := gomock.NewController(t)
	mockCreditScoreClient := mock_client.NewMockCreditScoreClient(mockCtrl)
	service := newLoanApprovalService(mockCreditScoreClient)

	// act
	ok, got := service.Sanction(age, float64(maxPossibleAmount))

	// assert
	if !ok {
		t.Error("false returned")
	}

	if want != got {
		t.Errorf("wanted %.2f, got %.2f", want, got)
	}
}
