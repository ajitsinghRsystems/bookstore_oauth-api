package access_token

import "errors"

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}
