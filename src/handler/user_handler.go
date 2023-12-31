package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/constant"
	"github.com/rxrav/loan_app/src/dto"
	"github.com/rxrav/loan_app/src/repo"
	"github.com/rxrav/loan_app/src/service"
)

func GetUserHandler(writer http.ResponseWriter, request *http.Request) {
	username := mux.Vars(request)["username"]
	log.Info().Msg(fmt.Sprintf("username received is %s", username))
	userService := service.UserServiceBuilder(service.DbUserRepoBuilder(repo.GetDbInstance(constant.PgConnectionString)))
	user := userService.GetUser(username)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(writer).Encode(user)
}

func CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	body, _ := io.ReadAll(request.Body)
	var newUser dto.CreateUserRequest
	_ = json.Unmarshal(body, &newUser)

	userService := service.UserServiceBuilder(service.DbUserRepoBuilder(repo.GetDbInstance(constant.PgConnectionString)))
	uuid, username := userService.CreateUser(newUser)
	formedResourceURL := constant.BaseUrl.JoinPath(strings.Replace(constant.GetUserRoute, "{username}", username, 1)).String()

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(writer).Encode(dto.CreateUserResponse{
		CreatedUUID:     uuid.String(),
		CreatedUsername: username,
		ResourceURL:     formedResourceURL,
	})
}