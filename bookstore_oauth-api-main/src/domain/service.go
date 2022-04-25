package access_token

import ("github.com/ajitsinghRsystems/bookstore_oauth-api-main/utils/errors"	
)

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}
