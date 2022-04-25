package app

import ("BookStore_OAuth-API-Main/repository/db"
		"BookStore_OAuth-API-Main/src/domain/accesstoken"
)


func StartApplication() {

	atService := accesstoken.NewService(db.NeRepository())

}