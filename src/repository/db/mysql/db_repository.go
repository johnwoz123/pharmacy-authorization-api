package mysql

import (
	"github.com/johnwoz123/pharmacy-authorization-api/src/domain/access_token"
	"github.com/johnwoz123/pharmacy-authorization-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErrors)
}

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}

// GetById database repository
func (r *repository) GetById(string) (*access_token.AccessToken, *errors.RestErrors) {
	return nil, errors.InternalServerError("database connection not implemented yet!")
}
