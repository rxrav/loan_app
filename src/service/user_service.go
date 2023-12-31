package service

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/domain"
	"github.com/rxrav/loan_app/src/dto"
	"github.com/rxrav/loan_app/src/repo"
)

type UserService interface {
	GetUser(username string) *dto.UserDetails
	CreateUser(user dto.CreateUserRequest) (uuid.UUID, string)
}

type DefaultUserService struct {
	userRepo repo.UserRepo
}

func randomSixDigitNumberGenerator() int {
	min := 100000
    max := 999999
    return rand.Intn(max - min) + min
}

func generateUsername(firstName, lastName string) string {
	return fmt.Sprintf("%s%s%d", strings.ToLower(firstName), strings.ToLower(lastName), randomSixDigitNumberGenerator())
}

func (s DefaultUserService) GetUser(username string) *dto.UserDetails {
	user := s.userRepo.GetUser(username)
	if user == nil {
		log.Error().Msg("user not found")
		panic(dto.LoanApplicationError{
			ErrCode:    10001,
			ErrDetails: "user not found",
		})
	}
	userDetails := &dto.UserDetails{
		UserID: user.ID.String(),
		Username: user.Username,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Age: user.Age,
		SocialNumber: user.SocialNumber,
	}
	return userDetails
}

func (s DefaultUserService) CreateUser(user dto.CreateUserRequest) (uuid.UUID, string) {
	_username := generateUsername(user.FirstName, user.LastName)
	userUUID := s.userRepo.CreateUser(domain.User{
		Username: _username,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Age: user.Age,
		SocialNumber: user.SocialNumber,
	})
	return userUUID, _username
}

func newUserService(_userRepo repo.UserRepo) DefaultUserService {
	return DefaultUserService{
		userRepo: _userRepo,
	}
}
