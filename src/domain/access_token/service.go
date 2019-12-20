package access_token

import (
	"github.com/johnwoz123/pharmacy-authorization-api/src/utils/errors"
	"strings"
)

// Service is responsible for handling the business logic for the domain

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErrors)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// GetById - business logic to retrieve information
func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErrors) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.BadRequestError("invalid access_token_id")
	}
	accessToken, err := s.repo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
