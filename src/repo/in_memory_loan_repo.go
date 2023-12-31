package repo

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/domain"
)

var inMemoryLoans []domain.Loan

type InMemoryLoanRepo struct{}

var inMemoryLoanRepo *InMemoryLoanRepo

func (r InMemoryLoanRepo) GetLoan(loanID uuid.UUID) *domain.Loan {
	for i := 0; i < len(inMemoryLoans); i ++ {
		if inMemoryLoans[i].ID == loanID {
			return &inMemoryLoans[i]
		}
	}
	return nil
}

func (r InMemoryLoanRepo) GetAllLoans(username string) []domain.Loan {
	var loans []domain.Loan
	for i := 0; i < len(inMemoryLoans); i ++ {
		if inMemoryLoans[i].Username == username {
			loans = append(loans, inMemoryLoans[i])
		}
	}
	return loans
}

func (r InMemoryLoanRepo) CreateLoan(loan domain.Loan) uuid.UUID {
	loan.ID = uuid.New()
	inMemoryLoans = append(inMemoryLoans, loan)
	return loan.ID
}

func GetInMemoryLoanRepoInstance() *InMemoryLoanRepo {
	if inMemoryLoanRepo == nil {
		log.Info().Msg("creating new singleton loan repo")
		loanRepoInitLock.Lock()
		defer loanRepoInitLock.Unlock()
		inMemoryLoanRepo = &InMemoryLoanRepo{}
	} else {
		log.Info().Msg("reusing singleton loan repo")
	}
	return inMemoryLoanRepo
}
