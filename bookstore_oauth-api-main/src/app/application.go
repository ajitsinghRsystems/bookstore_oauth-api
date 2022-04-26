package app

import (
	"BookStore_OAuth-API-Main/src/clients/cassandra"
	"BookStore_OAuth-API-Main/src/http"
	"BookStore_OAuth-API-Main/src/repository/db"
	"BookStore_OAuth-API-Main/src/services/access_token"

	"github.com/gin-gonic/gin"
)
var (
	router = gin.Default()
)
func StartApplication() {

	session, err:= cassandra.GetSession()
	if err != nil{
		panic(err)
	}
	session.Close()
	atService := access_token.NewService(db.NewRepository())
	atHandler:= http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id",atHandler.GetById)
	router.Run(":8080")

}