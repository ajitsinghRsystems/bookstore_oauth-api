package db

import (
	"BookStore_OAuth-API-Main/src/clients/cassandra"
	"BookStore_OAuth-API-Main/src/domain/access_token"
	"BookStore_OAuth-API-Main/src/utils/errors"
)
func NewRepository() DbRepository{
	return &dbRepositry{}
	
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}
type dbRepositry struct{

}

func (r *dbRepositry) GetById(id string) (*access_token.AccessToken, *errors.RestErr){
	session, err:= cassandra.GetSession()
	if err != nil{
		panic(err)
	}
	defer session.Close()

	return nil,errors.NewInternalServerError("Database connection not implemented")
}