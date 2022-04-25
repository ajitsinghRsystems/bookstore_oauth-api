package access_token

import "github.com/ajitsinghRsystems/bookstore_user-api/utils/errors"

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)

}
type service struct{

}
func NewService() Service{
	return &service{}
}
func (S *service) GetByID(string) (*AccessToken,*errors.RestErr)
{
	return nil,nil
}