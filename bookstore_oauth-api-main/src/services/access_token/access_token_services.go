package access_token

import (
	"BookStore_OAuth-API-Main/src/domain/access_token"
	"BookStore_OAuth-API-Main/src/repository/db"
	"BookStore_OAuth-API-Main/src/utils/errors"
)

type Service interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	
}

type Repository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type service struct {
	repository db.DbRepository
}

func NewService(repo db.DbRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}