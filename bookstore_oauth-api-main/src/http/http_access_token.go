package http

import (
	"BookStore_OAuth-API-Main/src/services/access_token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface{
	GetById( *gin.Context) 
}
type accessTokenHandler struct{
	service access_token.Service
}
func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
func (h *accessTokenHandler) GetById(c *gin.Context){
	c.JSON(http.StatusNotImplemented,"Implement me!")
}
