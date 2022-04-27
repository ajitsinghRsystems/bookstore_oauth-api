package db

import (
	"BookStore_OAuth-API-Main/src/clients/cassandra"
	"BookStore_OAuth-API-Main/src/domain/access_token"
	"BookStore_OAuth-API-Main/src/utils/errors"

	"github.com/gocql/gocql"
)
const(
	queryGetAccessToken ="SELECT access_token,user_id,client_id,expires from access_tokens where access_token=?;"
	queryCreateAccessToken="Insert into access_tokens ( access_token,user_id,client_id,expires) values (?,?,?,?);"
	queryUpdateAccessToken=" Update Access_tokens set expires=? where access_token=?;"
)
func NewRepository() DbRepository{
	return &dbRepositry{}
	
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create (access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime (access_token.AccessToken) *errors.RestErr
}
type dbRepositry struct{

}

func (r *dbRepositry) GetById(id string) (*access_token.AccessToken, *errors.RestErr){
	
	var result access_token.AccessToken
	if err:= cassandra.GetSession().Query(queryGetAccessToken,id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,); err !=nil{
			if err==gocql.ErrNotFound{
				return	nil, errors.NewNotFoundError("No access token found for the specified ID")
			}
			return nil,errors.NewInternalServerError(err.Error())
	}

	return &result,nil
}
func (r *dbRepositry) Create(at access_token.AccessToken) (*errors.RestErr){
	

	if err:= cassandra.GetSession().Query(queryCreateAccessToken,at.AccessToken,at.UserId,at.ClientId,at.Expires).Exec(); err !=nil{
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepositry) UpdateExpirationTime(at access_token.AccessToken) (*errors.RestErr){
	

	if err:= cassandra.GetSession().Query(queryUpdateAccessToken,at.Expires,at.AccessToken).Exec(); err !=nil{
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}