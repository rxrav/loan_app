package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/rxrav/loan_app/src/client"
	"github.com/rxrav/loan_app/src/repo"
)

func CreditScoreClientBuilder(baseUrl string) client.CreditScoreClient {
	return client.GetCreditScoreClientInstance(baseUrl)
}

func DbUserRepoBuilder(db *sqlx.DB) repo.UserRepo {
	return repo.GetDefaultUserRepoInstance(db)
}

func InMemUserRepoBuilder() repo.UserRepo {
	return repo.GetInMemoryUserRepoInstance()
}

func DbLoanRepoBuilder(db *sqlx.DB) repo.LoanRepo {
	return repo.GetDefaultLoanRepoInstance(db)
}

func InMemLoanRepoBuilder() repo.LoanRepo {
	return repo.GetInMemoryLoanRepoInstance()
}

func UserServiceBuilder(userRepo repo.UserRepo) UserService {
	return newUserService(userRepo)
}

func LoanApprovalServiceBuilder(defaultClient client.CreditScoreClient) LoanApprovalService {
	return newLoanApprovalService(defaultClient)
}

func LoanApplicationServiceBuilder(
	userRepo repo.UserRepo,
	defaultClient client.CreditScoreClient,
	db *sqlx.DB,
) LoanApplicationService {
	loanApplicationService := newLoanApplicationService(
		UserServiceBuilder(userRepo),
		LoanApprovalServiceBuilder(defaultClient),
		DbLoanRepoBuilder(db),
	)
	return loanApplicationService
}