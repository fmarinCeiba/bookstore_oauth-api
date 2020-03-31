package db

import (
	"github.com/fmarinCeiba/bookstore_oauth-api/domain/access_token"
	"github.com/fmarinCeiba/bookstore_oauth-api/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func (r *dbRepository) GetByID(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
