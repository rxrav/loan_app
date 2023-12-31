package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
)

var lock = &sync.Mutex{}

type CreditScoreClient interface {
	GetScore(userID string) (int, error)
}

type DefaultCreditScoreClient struct {
	baseUrl string
}

var client *DefaultCreditScoreClient

func (c DefaultCreditScoreClient) GetScore(socialNumber string) (int, error) {
	type respSchema struct {
		Score int
	}

	url := fmt.Sprintf("%sscore/%s", c.baseUrl, socialNumber)
	log.Info().Msg(url)

	apiResponse, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	if apiResponse.StatusCode == 404 {
		return 0, errors.New("unable to fetch credit score")
	}

	apiData, err := io.ReadAll(apiResponse.Body)
	if err != nil {
		return 0, err
	}

	var s respSchema
	_ = json.Unmarshal(apiData, &s)

	log.Info().Msg(fmt.Sprintf("Score is %d", s.Score))
	return s.Score, nil
}

func GetCreditScoreClientInstance(_baseUrl string) *DefaultCreditScoreClient {
	if client == nil {
		log.Info().Msg("creating new singleton client")
		lock.Lock()
		defer lock.Unlock()
		client = &DefaultCreditScoreClient{
			baseUrl: _baseUrl,
		}
	} else {
		log.Info().Msg("reusing singleton client")
	}
	return client
}
