package rest

import (
	"BookStore_OAuth-API-Main/src/domain/users"
	"BookStore_OAuth-API-Main/src/utils/errors"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
)
var(
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://localhost:8080",
		Timeout: 100 *time.Microsecond,
	}
)
type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}
type UsersRepository struct{}

func NewRepository() RestUserRepository {
	return &UsersRepository{}
}
func (r *UsersRepository) LoginUser( email string, password string) (*users.User, *errors.RestErr){
request := users.UserLoginRequest{
	Email: email,
	Password:password,
}
	response := usersRestClient.Post("/users/login",request)
	fmt.Println(response.String())

	if response ==nil || response.Response == nil{
		return nil,errors.NewInternalServerError("Invalide restclient response")
	}
	if response.StatusCode>299{
		var restErr errors.RestErr
		err:= json.Unmarshal(response.Bytes(),&restErr)
		if err !=nil{
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err:= json.Unmarshal(response.Bytes(),&user); err!=nil{
		return nil,errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user,nil
}