package access_token

import (
	"github.com/johnwoz123/pharmacy-authorization-api/src/utils/errors"
	"time"
)

const (
	expirationTime       = 24
	grantTypePassword    = "password"
	grantTypeClientCreds = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	// password grant type
	Username int64  `json:"user_name"`
	Password string `json:"password"`
	// used for client creds grant type
	ClientId     int64  `json:"client_id"`
	ClientSecret int64  `json:"client_secret"`
	Scope        string `json:"scope"`
}

func (token *AccessToken) Validate() *errors.RestErrors {
	if token.UserId <= 0 {
		return errors.BadRequestError("invalid user id")
	}

	if token.ClientId <= 0 {
		return errors.BadRequestError("invalid client id")
	}

	if token.Expires <= 0 {
		return errors.BadRequestError("invalid expiration time")
	}
	return nil
}

func GetAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) isTokenExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
