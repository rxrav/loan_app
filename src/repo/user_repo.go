package repo

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/domain"
	"github.com/rxrav/loan_app/src/dto"
)

type DefaultUserRepo struct {
	db *sqlx.DB
}

var defaultUserRepo *DefaultUserRepo

func (r DefaultUserRepo) GetUser(username string) *domain.User {
	query := `select id, user_name, first_name, last_name, user_age, social_number
		from loan_users
		where user_name = $1`

	var user domain.User
	err := r.db.Get(&user, query, username)

	if err != nil {
		log.Error().Msg(fmt.Sprintf("error while fetching user with username %s: Error %v", username, err))
		return nil
	}
	return &user
}

func (r DefaultUserRepo) CreateUser(user domain.User) uuid.UUID {
	user.ID = uuid.New()
	query := `insert into loan_users
		(id, user_name, first_name, last_name, user_age, social_number) 
		values
		($1, $2, $3, $4, $5, $6)`

	_, err := r.db.Exec(query, user.ID, user.Username, user.FirstName, user.LastName, user.Age, user.SocialNumber)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			errMsg := fmt.Sprintf("social number belongs to another user")
			log.Error().Msg(errMsg)
			panic(dto.LoanApplicationError{
				ErrCode:    20004,
				ErrDetails: errMsg,
			})
		} else {
			errMsg := fmt.Sprintf("error while creating user with username %s: Error: %v", user.Username, err)
			log.Error().Msg(errMsg)
			panic(dto.LoanApplicationError{
				ErrCode:    20002,
				ErrDetails: errMsg,
			})
		}
	}
	return user.ID
}

func GetDefaultUserRepoInstance(_db *sqlx.DB) *DefaultUserRepo {
	if defaultUserRepo == nil {
		log.Info().Msg("creating new singleton user repo")
		userRepoInitLock.Lock()
		defer userRepoInitLock.Unlock()
		defaultUserRepo = &DefaultUserRepo{
			db: _db,
		}
	} else {
		log.Info().Msg("reusing singleton user repo")
	}
	return defaultUserRepo
}
