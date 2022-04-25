package db

import (
	"BookStore_OAuth-API-Main/utils/errors"
	"BookStore_OAuth-API-Main/src/domain/access_token"
)

func NeRepository() DbRepository{
return &dbRepository{}
}
type DbRepository interface {
	GetByID(string) ( *access_token.AccessToken, *errors.RestErr)
}
type dbRepository struct{}

func(r *dbRepository) GetByID(id string) (*access_token.AccessToken,*errors.RestErr){
	return nil,nil
}
