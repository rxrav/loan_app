package repo

import (
	"sync"

	"github.com/google/uuid"
	"github.com/rxrav/loan_app/src/domain"
)

var userRepoInitLock = &sync.Mutex{}
var loanRepoInitLock = &sync.Mutex{}

type UserRepo interface {
	GetUser(username string) *domain.User
	CreateUser(user domain.User) uuid.UUID
}

type LoanRepo interface {
	GetLoan(loanID uuid.UUID) *domain.Loan
	GetAllLoans(username string) []domain.Loan
	CreateLoan(loan domain.Loan) uuid.UUID
}
