package access_token

import (
	"github.com/johnwoz123/pharmacy-authorization-api/src/utils/errors"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErrors)
}

func GetAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) isTokenExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
