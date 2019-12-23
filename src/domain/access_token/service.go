package access_token

import (
	"github.com/johnwoz123/pharmacy-authorization-api/src/utils/crypto"
	"github.com/johnwoz123/pharmacy-authorization-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErrors)
	Create(AccessToken) *errors.RestErrors
	UpdateExpirationTime(AccessToken) *errors.RestErrors
}

// Service is responsible for handling the business logic for the domain

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErrors)
	Create(AccessToken) *errors.RestErrors
	UpdateExpirationTime(AccessToken) *errors.RestErrors
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

func (s *service) Create(token AccessToken) *errors.RestErrors {
	if err := token.Validate(); err != nil {
		return err
	}
	token.AccessToken = crypto.GenerateToken()
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if len(token.AccessToken) == 0 {
		return errors.BadRequestError("invalid access_token_id")
	}
	return s.repo.Create(token)
}

func (s *service) UpdateExpirationTime(time AccessToken) *errors.RestErrors {
	if err := time.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateExpirationTime(time)
}
