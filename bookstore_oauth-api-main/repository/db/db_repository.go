package db

import (
	"BookStore_OAuth-API-Main/src/domain/accesstoken"
	"BookStore_OAuth-API-Main/utils/errors"
)

func NeRepository() DbRepository{
return &dbRepository{}
}
type DbRepository interface {
	GetByID(string) ( *accesstoken.AccessToken, *errors.RestErr)
}
type dbRepository struct{}

func(r *dbRepository) GetByID(id string) (*accesstoken.AccessToken,*errors.RestErr){
	return nil,nil
}
