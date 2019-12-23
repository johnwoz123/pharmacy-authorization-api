package rest

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("starting test cases......")
	//rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeout(t *testing.T) {
	//mock object
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://127.0.01:8081/login",
		ReqBody:      `{"email":"test@me.com","password":"Test123!"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repo := usersDbRestRepo{}
	user, err := repo.LoginUser("test@me.com", "Test123!")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when logging in", err.Message)
}

func TestLoginUserInvalidRestErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidCredentials(t *testing.T) {

}

func TestLoginUserUnmarshallingError(t *testing.T) {

}

func TestInvalidRestResponse(t *testing.T) {

}
