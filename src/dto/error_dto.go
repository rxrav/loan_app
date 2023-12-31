package dto

import "fmt"

type LoanApplicationError struct {
	ErrCode    int    `json:"error_code"`
	ErrDetails string `json:"error_details"`
}

func (e *LoanApplicationError) Error() string {
	return fmt.Sprintf("err code [%d], Details: %v", e.ErrCode, e.ErrDetails)
}