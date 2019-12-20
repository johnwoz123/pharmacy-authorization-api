package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime)
}

func TestGetAccessToken(t *testing.T) {
	at := GetAccessToken()
	assert.False(t, at.isTokenExpired(), "new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "access token should not have defined access token id")
	assert.True(t, at.UserId == 0, "access token should not have associated user id")
}

func TestExpirationOfAccessToken(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.isTokenExpired(), "empty access token should be empty by default")

	at.Expires = time.Now().UTC().Add(5 * time.Hour).Unix()

	assert.False(t, at.isTokenExpired(), "access token that expires 3 hours from now should not be expired")
}
