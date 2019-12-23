package rest

import (
	"encoding/json"
	"fmt"
	"github.com/johnwoz123/pharmacy-authorization-api/src/domain/users"
	"github.com/johnwoz123/pharmacy-authorization-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	restClient = rest.RequestBuilder{
		Headers: nil,
		Timeout: 100 * time.Millisecond,
		BaseURL: "http://localhost:8081",
	}
)

type UserRestRepo interface {
	LoginUser(string, string) (*users.User, *errors.RestErrors)
}

type usersDbRestRepo struct{}

func NewRepo() UserRestRepo {
	return &usersDbRestRepo{}
}

func (u usersDbRestRepo) LoginUser(email string, password string) (*users.User, *errors.RestErrors) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	bytes, _ := json.Marshal(req)
	fmt.Println(string(bytes))
	response := restClient.Post("/login", req)
	b, _ := json.Marshal(response)
	fmt.Println(string(b))
	fmt.Println(response)
	if response == nil || response.Response == nil {
		return nil, errors.InternalServerError("invalid rest client response when logging in")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErrors
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.InternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.InternalServerError("error unmarshaling users response")
	}
	return &user, nil
}
