package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/domain"
	"github.com/rxrav/loan_app/src/dto"
	"github.com/rxrav/loan_app/src/repo"
)

type LoanApplicationService interface {
	Apply(request dto.LoanApplicationRequest) dto.LoanApplicationResponse
	GetLoan(loanID string) *domain.Loan
	GetAllLoans(username string) []domain.Loan
}

type DefaultLoanApplicationService struct {
	userService         UserService
	loanRepo            repo.LoanRepo
	loanApprovalService LoanApprovalService
}

func (s DefaultLoanApplicationService) Apply(request dto.LoanApplicationRequest) dto.LoanApplicationResponse {
	if request.AppliedAmount > 100_000 {
		log.Error().Msg("loan application amount is too high")
		panic(dto.LoanApplicationError{
			ErrCode:    10004,
			ErrDetails: "loan application amount is too high",
		})
	}

	user := s.userService.GetUser(request.Username)
	response := dto.LoanApplicationResponse{}
	maxPossibleAmount := s.loanApprovalService.Validate(user.SocialNumber, request.AppliedAmount)

	// var message string
	if isApproved, approvedAmount := s.loanApprovalService.Sanction(user.Age, maxPossibleAmount); isApproved {
		loanID := s.loanRepo.CreateLoan(domain.Loan{
			Username:       request.Username,
			AppliedAmount:  request.AppliedAmount,
			ApprovedAmount: approvedAmount,
			Status:         "Approved",
		})
		response.AppliedAmount = request.AppliedAmount
		response.ApprovedAmount = approvedAmount
		response.IsApproved = true
		response.Message = fmt.Sprintf("Your loan was appoved. Approved amount is USD %.2f", approvedAmount)
		response.Username = request.Username
		response.LoanRequestID = loanID.String()
	} else {
		loanID := s.loanRepo.CreateLoan(domain.Loan{
			ID:             uuid.New(),
			Username:       request.Username,
			AppliedAmount:  request.AppliedAmount,
			ApprovedAmount: approvedAmount,
			Status:         "Rejected",
		})
		response.AppliedAmount = request.AppliedAmount
		response.ApprovedAmount = 0
		response.IsApproved = false
		response.Message = "Your loan was not approved"
		response.Username = request.Username
		response.LoanRequestID = loanID.String()
	}
	return response
}

func (s DefaultLoanApplicationService) GetLoan(loanID string) *domain.Loan {
	loanUUID, err := uuid.Parse(loanID)
	if err != nil {
		panic(dto.LoanApplicationError{
			ErrCode:    10002,
			ErrDetails: "not a valid loan ID",
		})
	}
	loan := s.loanRepo.GetLoan(loanUUID)
	if loan == nil {
		panic(dto.LoanApplicationError{
			ErrCode:    10003,
			ErrDetails: "loan not found",
		})
	}
	return loan
}

func (s DefaultLoanApplicationService) GetAllLoans(username string) []domain.Loan {
	return s.loanRepo.GetAllLoans(username)
}

func newLoanApplicationService(
	_userService UserService,
	_loanApprovalService LoanApprovalService,
	_loanRepo repo.LoanRepo) DefaultLoanApplicationService {
	log.Info().Msg("creating new loan application service")
	return DefaultLoanApplicationService{
		userService:         _userService,
		loanApprovalService: _loanApprovalService,
		loanRepo:            _loanRepo,
	}
}
