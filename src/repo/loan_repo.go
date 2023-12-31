package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/domain"
	"github.com/rxrav/loan_app/src/dto"
)

type DefaultLoanRepo struct {
	db *sqlx.DB
}

var defaultLoanRepo *DefaultLoanRepo

func (r DefaultLoanRepo) GetLoan(loanID uuid.UUID) *domain.Loan {
	query := `select 
		id, 
		user_name, 
		applied_amount, 
		approved_amount, 
		status_ 
		from loan_applications where id = $1`

	var loan domain.Loan
	err := r.db.Get(&loan, query, loanID)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("error while fetching loan with id %s: Error %v", loanID.String(), err))
		return nil
	}
	return &loan
}

func (r DefaultLoanRepo) GetAllLoans(username string) []domain.Loan {
	query := `select 
		id, 
		user_name, 
		applied_amount, 
		approved_amount, 
		status_ 
		from loan_applications where user_name = $1`

	var loan []domain.Loan
	err := r.db.Select(&loan, query, username)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("error while fetching loan for username %s: Error %v", username, err))
		return []domain.Loan{}
	}
	return loan
}

func (r DefaultLoanRepo) CreateLoan(loan domain.Loan) uuid.UUID {
	loan.ID = uuid.New()
	query := `insert into loan_applications
		(id, user_name, applied_amount, approved_amount, status_) 
		values
		($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query, loan.ID, loan.Username, loan.AppliedAmount, loan.ApprovedAmount, loan.Status)
	if err != nil {
		errMsg := fmt.Sprintf("error while creating loan application for username %s: Error: %v", loan.Username, err)
		log.Error().Msg(errMsg)
		panic(dto.LoanApplicationError{
			ErrCode: 20001,
			ErrDetails: errMsg,
		})
	}
	return loan.ID
}

func GetDefaultLoanRepoInstance(_db *sqlx.DB) *DefaultLoanRepo {
	if inMemoryLoanRepo == nil {
		log.Info().Msg("creating new singleton loan repo")
		loanRepoInitLock.Lock()
		defer loanRepoInitLock.Unlock()
		defaultLoanRepo = &DefaultLoanRepo{
			db: _db,
		}
	} else {
		log.Info().Msg("reusing singleton loan repo")
	}
	return defaultLoanRepo
}
