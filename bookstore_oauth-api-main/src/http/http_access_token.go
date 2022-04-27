package http

import (
	"BookStore_OAuth-API-Main/src/services/access_token"
	"net/http"
	"strings"
	atDomain "BookStore_OAuth-API-Main/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"BookStore_OAuth-API-Main/src/utils/errors"
)

type AccessTokenHandler interface{
	GetById( *gin.Context) 
	Create(*gin.Context)
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
func (h *accessTokenHandler) Create(c *gin.Context){
	var request atDomain.AccessToken
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	 err := h.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, request)
}