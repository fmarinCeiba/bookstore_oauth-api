package access_token

import (
	"strings"

	"github.com/fmarinCeiba/bookstore_oauth-api/domain/access_token"
	"github.com/fmarinCeiba/bookstore_oauth-api/repository/db"
	"github.com/fmarinCeiba/bookstore_oauth-api/repository/rest"
	"github.com/fmarinCeiba/bookstore_oauth-api/utils/errors"
)

type Service interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

func NewService(usersRepo rest.RestUsersRepository, repoDB db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        repoDB,
	}
}

func (s *service) GetByID(accessTokenID string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	aToken, err := s.dbRepo.GetByID(accessTokenID)

	if err != nil {
		return nil, err
	}
	return aToken, nil
}

func (s *service) Create(atr access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := atr.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	//Authenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(atr.Username, atr.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
