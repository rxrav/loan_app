package integrationtests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rxrav/loan_app/src/dto"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rxrav/loan_app/src/client"
	"github.com/rxrav/loan_app/src/domain"
	"github.com/rxrav/loan_app/src/handler"
	"github.com/rxrav/loan_app/src/repo"
	"github.com/rxrav/loan_app/src/service"
)

func Test_Integration_GetLoanHandler(t *testing.T) {
	// arrange
	request := httptest.NewRequest(http.MethodGet, "/loan", nil)
	request = mux.SetURLVars(request, map[string]string{"loanID": "1da322d1-84ec-4f23-976d-e3a8512f46a9"})
	writer := httptest.NewRecorder()

	loanApplicationService := service.LoanApplicationServiceBuilder(
		service.DbUserRepoBuilder(repo.GetDbInstance(tcConnStr)),
		client.GetCreditScoreClientInstance(fmt.Sprintf("http://localhost:%s/", csPort.Port())),
		repo.GetDbInstance(tcConnStr),
	)

	// act
	handler.GetLoanHandler(writer, request, loanApplicationService)

	// assert
	res := writer.Result()
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	var loan domain.Loan
	_ = json.Unmarshal(data, &loan)
	if res.StatusCode != 200 {
		t.Error("no match")
	}
	if loan.ID.String() != "1da322d1-84ec-4f23-976d-e3a8512f46a9" {
		t.Error("not that loan")
	}
}

func Test_Integration_ApplyLoanHandler(t *testing.T) {
	requestDto := dto.LoanApplicationRequest{
		Username:      "janedoe444444",
		AppliedAmount: 20000,
	}
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(requestDto)
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodPost, "/apply", &b)
	writer := httptest.NewRecorder()
	loanApplicationService := service.LoanApplicationServiceBuilder(
		service.DbUserRepoBuilder(repo.GetDbInstance(tcConnStr)),
		client.GetCreditScoreClientInstance(fmt.Sprintf("http://localhost:%s/", csPort.Port())),
		repo.GetDbInstance(tcConnStr),
	)
	// act
	handler.ApplyLoanHandler(writer, request, loanApplicationService)

	// assert
	res := writer.Result()
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	var responseDto dto.LoanApplicationResponse
	_ = json.Unmarshal(data, &responseDto)

	if res.StatusCode != 202 {
		t.Error("not accepted")
	}
	if responseDto.Username != "janedoe444444" {
		t.Error("user name is not janedoe444444")
	}
	if responseDto.ApprovedAmount != 18000 {
		t.Error("approved amount is not 18000")
	}
	if responseDto.IsApproved != true {
		t.Error("is approved is false")
	}
}

func Test_Integration_GetAllLoansHandler(t *testing.T) {
	// arrange
	request := httptest.NewRequest(http.MethodGet, "/loan/all?username=brucewayne333333", nil)
	writer := httptest.NewRecorder()

	loanApplicationService := service.LoanApplicationServiceBuilder(
		service.DbUserRepoBuilder(repo.GetDbInstance(tcConnStr)),
		client.GetCreditScoreClientInstance(fmt.Sprintf("http://localhost:%s/", csPort.Port())),
		repo.GetDbInstance(tcConnStr),
	)

	// act
	handler.GetAllLoansHandler(writer, request, loanApplicationService)

	// assert
	res := writer.Result()
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	var loans []domain.Loan
	_ = json.Unmarshal(data, &loans)

	if res.StatusCode != 200 {
		t.Error("no match")
	}
	if len(loans) != 1 {
		t.Error("not one loan")
	}
}
