package accesstoken

import "BookStore_OAuth-API-Main/utils/errors"

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Repository interface{
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct{
repository Repository

}
func NewService( repo Repository) Service{
return &service{
	repository : repo,
}
}

func (s *service) GetByID(string) (*AccessToken, *errors.RestErr){
	return nil,nil
}