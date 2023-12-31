package dto

type LoanApplicationRequest struct {
	Username      string  `json:"username"`
	AppliedAmount float64 `json:"applied_amount"`
}

type LoanApplicationResponse struct {
	Username       string  `json:"username"`
	AppliedAmount  float64 `json:"applied_amount"`
	IsApproved     bool    `json:"is_approved"`
	ApprovedAmount float64 `json:"approved_amount"`
	Message        string  `json:"message"`
	LoanRequestID  string  `json:"loan_request_id"`
}

