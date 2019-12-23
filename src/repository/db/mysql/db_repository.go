package mysql

import (
	"github.com/johnwoz123/pharmacy-authorization-api/src/clients/mysql"
	"github.com/johnwoz123/pharmacy-authorization-api/src/domain/access_token"
	"github.com/johnwoz123/pharmacy-authorization-api/src/utils/errors"
)

const (
	getAccessTokenQuery    = "SELECT * FROM access_token where access_token = ?"
	insertAccessTokenQuery = "INSERT INTO access_token(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	updateAccessTokenQuery = "UPDATE access_token set expires = ?, WHERE access_token = ?"
)

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErrors)
	Create(access_token.AccessToken) *errors.RestErrors
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErrors
}

type repository struct {
}

func NewRepo() Repository {
	return &repository{}
}

// GetById database repository
func (r *repository) GetById(id string) (*access_token.AccessToken, *errors.RestErrors) {
	// TODO: get access token from Database Implementation
	stmtFindById, err := mysql.Client.Prepare(getAccessTokenQuery)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer stmtFindById.Close()

	row := mysql.Client.QueryRow(getAccessTokenQuery, id)

	var accessToken access_token.AccessToken
	rowErr := row.Scan(&accessToken.AccessToken, &accessToken.UserId, &accessToken.ClientId, &accessToken.Expires)
	if rowErr != nil {
		if rowErr.Error() == "sql: no rows in result set" {
			return nil, errors.InternalServerError("no access token found")
		}
		return nil, errors.InternalServerError(rowErr.Error())
	}

	return &accessToken, nil
}

func (r *repository) Create(token access_token.AccessToken) *errors.RestErrors {
	stmtInsert, err := mysql.Client.Prepare(insertAccessTokenQuery)
	if err != nil {
		return errors.InternalServerError("error with the database")
	}
	// since there are no errors defer close until complete
	defer stmtInsert.Close()
	_, insertErr := stmtInsert.Exec(token.AccessToken, token.UserId, token.ClientId, token.Expires)
	if insertErr != nil {
		return errors.InternalServerError("error with the database")
	}
	return nil
}

func (r *repository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErrors {
	stmtUpdate, err := mysql.Client.Prepare(updateAccessTokenQuery)
	if err != nil {
		return errors.InternalServerError("error with the database")
	}
	// since there are no errors defer close until complete
	defer stmtUpdate.Close()
	_, err = stmtUpdate.Exec(token.Expires, token.AccessToken)

	if err != nil {
		return errors.InternalServerError("error updating the access token")
	}

	return nil

}
