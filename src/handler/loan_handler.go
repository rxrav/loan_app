package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/dto"
	"github.com/rxrav/loan_app/src/service"
)

func GetLoanHandler(
	writer http.ResponseWriter, 
	request *http.Request,
	loanApplicationService service.LoanApplicationService) {
	loanID := mux.Vars(request)["loanID"]
	
	log.Info().Msg(fmt.Sprintf("loanID received is %s", loanID))
	loan := loanApplicationService.GetLoan(loanID)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(loan)
}

func ApplyLoanHandler(
	writer http.ResponseWriter, 
	request *http.Request,
	loanApplicationService service.LoanApplicationService) {
	body, _ := io.ReadAll(request.Body)
	var loanRequest dto.LoanApplicationRequest
	_ = json.Unmarshal(body, &loanRequest)

	log.Info().Msg(fmt.Sprintf("payload %v", loanRequest))
	loanResponse := loanApplicationService.Apply(loanRequest)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(writer).Encode(loanResponse)
}

func GetAllLoansHandler(
	writer http.ResponseWriter, 
	request *http.Request, 
	loanApplicationService service.LoanApplicationService) {
	username := request.URL.Query().Get("username")
	if len(username) == 0 {
		panic(dto.LoanApplicationError{
			ErrCode:    9001,
			ErrDetails: "need to provide a username",
		})
	}
	log.Info().Msg(fmt.Sprintf("username received is %s", username))
	allLoans := loanApplicationService.GetAllLoans(username)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(allLoans)
}