package access_token

import (
	"BookStore_OAuth-API-Main/src/utils/errors"
	"strings"
	"time"

	)
const(
	expirationTime=24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"ClientId"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken)Validate ()  *errors.RestErr{
	at.AccessToken= strings.TrimSpace(at.AccessToken)
	if (at.AccessToken)==""{
		return  errors.NewBadRequestError("Invald access token")
	}
	if at.UserId <=0{
		return  errors.NewBadRequestError("InvaldUserId")	
	}
	if at.ClientId<=0{
		return  errors.NewBadRequestError("Invald Client Id")	
	}
	if at.Expires<=0{
		return  errors.NewBadRequestError("Invald Expire")	
	}

	return nil
}
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}
func (at AccessToken) IsExpired() bool{
	return time.Unix(at.Expires,0).Before(time.Now().UTC())
}