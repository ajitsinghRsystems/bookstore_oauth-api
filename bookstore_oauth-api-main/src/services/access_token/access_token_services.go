package access_token

import (
	"BookStore_OAuth-API-Main/src/domain/access_token"
	"BookStore_OAuth-API-Main/src/repository/db"
	"BookStore_OAuth-API-Main/src/utils/errors"
	
)

type Service interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
	
}


type service struct {
	repository db.DbRepository
}

func NewService(repo db.DbRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err!=nil{
		return nil,err
	}
	return accessToken, nil
}
func (s *service) Create(at access_token.AccessToken) *errors.RestErr{
	if err := at.Validate(); err != nil{
		return  err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr{
	if err := at.Validate(); err != nil{
		return  err
	}
	return s.repository.UpdateExpirationTime(at)
}