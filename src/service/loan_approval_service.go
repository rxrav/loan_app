package service

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/client"
	"github.com/rxrav/loan_app/src/dto"
)

type LoanApprovalService interface {
	Validate(socialNumber string, appliedAmount float64) float64
	Sanction(age int, maxPossibleAmount float64) (bool, float64)
}

type DefaultLoanApprovalService struct {
	creditScoreClient client.CreditScoreClient
}

func (s DefaultLoanApprovalService) Validate(socialNumber string, appliedAmount float64) float64 {
	maxPossibleAmount := 0.0
	score, err := s.creditScoreClient.GetScore(socialNumber)
	if err != nil {
		errMsg := fmt.Sprintf("unable to fetch credit score for socialNumber %s. Error: %v", socialNumber, err)
		panic(dto.LoanApplicationError{
			ErrCode:    9002,
			ErrDetails: errMsg,
		})
	}

	// score 800+ 		maxPossibleAmount = 90%
	// score 700 - 800	maxPossibleAmount = 80%
	// score 600 - 700	maxPossibleAmount = 70%
	// score 500 - 600	maxPossibleAmount = 50%
	// score below 500	maxPossibleAmount = 20%
	switch {
	case score > 800:
		maxPossibleAmount = float64(appliedAmount) * 0.9
	case score > 700 && score <= 800:
		maxPossibleAmount = float64(appliedAmount) * 0.8
	case score > 600 && score <= 700:
		maxPossibleAmount = float64(appliedAmount) * 0.7
	case score > 500 && score <= 600:
		maxPossibleAmount = float64(appliedAmount) * 0.6
	case score > 400 && score <= 500:
		maxPossibleAmount = float64(appliedAmount) * 0.5
	default:
		maxPossibleAmount = float64(appliedAmount) * 0.2
	}
	return maxPossibleAmount
}

func (s DefaultLoanApprovalService) Sanction(age int, maxPossibleAmount float64) (bool, float64) {
	approvedAmount := 0.0

	if age >= 50 {
		approvedAmount = maxPossibleAmount * 0.6
	} else if age >= 30 && age < 50 {
		approvedAmount = maxPossibleAmount * 0.8
	} else if age >= 18 && age < 30 {
		approvedAmount = maxPossibleAmount
	} else {
		return false, approvedAmount
	}
	return true, approvedAmount
}

func newLoanApprovalService(_creditScoreClient client.CreditScoreClient) DefaultLoanApprovalService {
	log.Info().Msg("creating new loan approval service")
	return DefaultLoanApprovalService{
		creditScoreClient: _creditScoreClient,
	}
}
