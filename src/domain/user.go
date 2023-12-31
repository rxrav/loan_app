package domain

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `db:"id"`
	Username     string    `db:"user_name"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	Age          int       `db:"user_age"`
	SocialNumber string    `db:"social_number"`
}

