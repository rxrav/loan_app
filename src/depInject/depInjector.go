package depInject

import (
	"github.com/rxrav/loan_app/src/constant"
	"github.com/rxrav/loan_app/src/repo"
	"github.com/rxrav/loan_app/src/service"
)

// CreateLoanApplicationService building the loan application service DAG here by injecting all other dependencies
// required to build this service
func CreateLoanApplicationService() service.LoanApplicationService {
	return service.LoanApplicationServiceBuilder(
		service.DbUserRepoBuilder(repo.GetDbInstance(constant.PgConnectionString)),
		service.CreditScoreClientBuilder(constant.CredScoreClientBaseUrl),
		repo.GetDbInstance(constant.PgConnectionString),
	)
}

// CreateUserService building the user service DAG here by injecting all other dependencies
// required to build this service
func CreateUserService() service.UserService {
	// Use this - Coding exercise
	return service.UserServiceBuilder(
		service.DbUserRepoBuilder(repo.GetDbInstance(constant.PgConnectionString)),
	)
}
