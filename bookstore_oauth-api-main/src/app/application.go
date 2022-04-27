package app

import (
	"BookStore_OAuth-API-Main/src/http"
	"BookStore_OAuth-API-Main/src/repository/db"
	"BookStore_OAuth-API-Main/src/services/access_token"

	"github.com/gin-gonic/gin"
)
var (
	router = gin.Default()
)
func StartApplication() {

	atService := access_token.NewService(db.NewRepository())
	atHandler:= http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id",atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")

}