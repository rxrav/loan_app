package domain

import "github.com/google/uuid"

type Loan struct {
	ID             uuid.UUID `db:"id"`
	Username       string    `db:"user_name"`
	AppliedAmount  float64   `db:"applied_amount"`
	ApprovedAmount float64   `db:"approved_amount"`
	Status         string    `db:"status_"`
}
