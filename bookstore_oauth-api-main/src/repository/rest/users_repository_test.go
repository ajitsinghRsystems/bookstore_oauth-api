package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}
func TestLoginUserTimeoutFromAPI(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL: "https://localhost:8080/users/login",
		HTTPMethod: http.MethodPost,
		ReqBody: `{"email":"email@gmail.com","password":"test123"}`,
		RespHTTPCode: -1,
		RespBody: `{}`,
})
	repository :=UsersRepository{}
user,err:= repository.LoginUser("email@gmail.com","test123")
assert.Nil(t,user)
assert.NotNil(t,err)
assert.EqualValues(t,http.StatusInternalServerError,err.Status)
assert.EqualValues(t,"Invalide restclient response",err.Message)
}
func TestLoginUserInvalidErrorInterface(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://localhost:8080/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
	})

	repository := UsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}
func TestLoginUserInvalidLoginCredential(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://localhost:8080/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "invalid login credentials", "status": 404, "error": "not_found"}`,
	})

	repository :=  UsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}
func TestLoginUserInvalidUserJsonResponse(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://localhost:8080/users/login",
		ReqBody:      `{"email":"email@gmail.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": "Fede", "last_name": "León", "email": "fedeleon.cba@gmail.com"}`,
	})

	repository := UsersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
//	assert.EqualValues(t, "error when trying to unmarshal users login response", err.Message)
}
func TestLoginUserNoError(t *testing.T){
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://localhost:8081/users/login",
		ReqBody:      `{"email":"zakio.tomar@rsystems.com","password":"Test345"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 5, "first_name": "Zakio", "last_name": "Tomar", "email": "zakio.tomar@rsystems.com","date_created":"","status":"Active"}`,
	})

	repository := UsersRepository{}

	user, err := repository.LoginUser("zakio.tomar@rsystems.com", "Test345")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 5, user.Id)
	assert.EqualValues(t, "Zakio", user.FirstName)
	assert.EqualValues(t, "Tomar", user.LastName)
	assert.EqualValues(t, "zakio.tomar@rsystems.com", user.Email)
}