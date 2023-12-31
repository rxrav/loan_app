package repo

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/domain"
)

var inMemoryUsers = map[string]domain.User{
	"johndoe111111": {
		ID:           uuid.New(),
		Username:     "johndoe111111",
		FirstName:    "John",
		LastName:     "Doe",
		Age:          35,
		SocialNumber: "AAA12001",
	},
	"peterparker222222": {
		ID:           uuid.New(),
		Username:     "peterparker222222",
		FirstName:    "Peter",
		LastName:     "Parker",
		Age:          16,
		SocialNumber: "AAA98121",
	},
	"brucewayne333333": {
		ID:           uuid.New(),
		Username:     "brucewayne333333",
		FirstName:    "Bruce",
		LastName:     "Wayne",
		Age:          52,
		SocialNumber: "AAA65091",
	},
	"janedoe444444": {
		ID:           uuid.New(),
		Username:     "janedoe444444",
		FirstName:    "Jane",
		LastName:     "Doe",
		Age:          26,
		SocialNumber: "AAA78034",
	},
	"clarkkent555555": {
		ID:           uuid.New(),
		Username:     "clarkkent555555",
		FirstName:    "Clark",
		LastName:     "Kent",
		Age:          52,
		SocialNumber: "AAA79811",
	},
}

type InMemoryUserRepo struct{}

var inMemoryUserRepo *InMemoryUserRepo

func (r InMemoryUserRepo) GetUser(username string) *domain.User {
	user, ok := inMemoryUsers[username]
	if ok {
		return &user
	}
	return nil
}

func (r InMemoryUserRepo) CreateUser(user domain.User) uuid.UUID {
	user.ID = uuid.New()
	inMemoryUsers[user.Username] = user
	return user.ID
}

func GetInMemoryUserRepoInstance() *InMemoryUserRepo {
	if inMemoryUserRepo == nil {
		log.Info().Msg("creating new singleton user repo")
		userRepoInitLock.Lock()
		defer userRepoInitLock.Unlock()
		inMemoryUserRepo = &InMemoryUserRepo{}
	} else {
		log.Info().Msg("reusing singleton user repo")
	}
	return inMemoryUserRepo
}
