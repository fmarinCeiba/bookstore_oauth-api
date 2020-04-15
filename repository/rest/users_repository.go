package rest

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/fmarinCeiba/bookstore_oauth-api/domain/users"
	"github.com/fmarinCeiba/bookstore_utils-go/rest_errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Microsecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (ur *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	resp := usersRestClient.Post("/users/login", req)
	if resp == nil || resp.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid resclient response when trying to login user", errors.New("rest error"))
	}
	if resp.StatusCode > 299 {
		var rErr rest_errors.RestErr
		err := json.Unmarshal(resp.Bytes(), &rErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when to trying to login user", errors.New("rest error"))
		}
		return nil, rErr
	}
	var u users.User
	if err := json.Unmarshal(resp.Bytes(), &u); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users login response", errors.New("rest error"))
	}
	return &u, nil
}
