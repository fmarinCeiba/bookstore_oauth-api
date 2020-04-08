package rest

import (
	"encoding/json"
	"time"

	"github.com/fmarinCeiba/bookstore_oauth-api/domain/users"
	"github.com/fmarinCeiba/bookstore_oauth-api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 100 * time.Microsecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRepository() RestUserRepository {
	return &usersRepository{}
}

func (ur *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	resp := usersRestClient.Post("/users/login", req)
	if resp == nil || resp.Response == nil {
		return nil, errors.NewInternalServerError("invalid resclient response when trying to login user")
	}
	if resp.StatusCode > 299 {
		var rErr errors.RestErr
		err := json.Unmarshal(resp.Bytes(), &rErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when to trying to login user")
		}
		return nil, &rErr
	}
	var u users.User
	if err := json.Unmarshal(resp.Bytes(), &u); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &u, nil
}
