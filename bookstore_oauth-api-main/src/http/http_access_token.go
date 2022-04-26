package http

import (
	"BookStore_OAuth-API-Main/src/services/access_token"
	"net/http"
	"strings"

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
	accessTokenId:=strings.TrimSpace(c.Param("access_token_id"))
	accessToken,err :=h.service.GetByID(accessTokenId)
	if err!=nil{
		c.JSON(err.Status,err)
		return
	}

	c.JSON(http.StatusOK,accessToken)
}
