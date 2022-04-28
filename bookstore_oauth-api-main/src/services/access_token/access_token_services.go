package access_token

import (
	"BookStore_OAuth-API-Main/src/domain/access_token"
	"BookStore_OAuth-API-Main/src/repository/db"
	"BookStore_OAuth-API-Main/src/repository/rest"
	"BookStore_OAuth-API-Main/src/utils/errors"
)

type Service interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
	
}


type service struct {
	restUserRepo rest.RestUserRepository
	repository db.DbRepository
}

func NewService(userRepo rest.RestUserRepository,repo db.DbRepository) Service {
	return &service{
		restUserRepo: userRepo,
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
func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken,*errors.RestErr){
	
	if err := request.Validate(); err != nil{
		return  nil,err
	}
	
	user, err := s.restUserRepo.LoginUser(request.Username,request.Password)

	if err !=nil{
		return nil,err
	}
	at:=access_token.GetNewAccessToken(user.Id)
	at.Generate()
	if err:=s.repository.Create(at); err!=nil{
		return nil,err
	}

	return &at,nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr{
	if err := at.Validate(); err != nil{
		return  err
	}
	return s.repository.UpdateExpirationTime(at)
}