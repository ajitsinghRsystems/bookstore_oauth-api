package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestAccessTokenConstants(t *testing.T){
	assert.EqualValues(t,24,expirationTime,"Expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T){
	at :=GetNewAccessToken()
	assert.False(t,at.IsExpired(),"Brand new access token should not be expired")
	assert.EqualValues(t,"",at.AccessToken,"New access token should not be empty")
	assert.True(t,at.UserId==0,"New access token should not associated with user")
}
func TestAccessTokenIsExpired(t *testing.T){
	at :=AccessToken{}
	assert.True(t, at.IsExpired(),"empty access token should be expired by default")	
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t,at.IsExpired(),"access token expiring three hours from now")

}