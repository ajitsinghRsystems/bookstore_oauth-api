package app

import ("BookStore_OAuth-API-Main/repository/db"
		"BookStore_OAuth-API-Main/src/domain/access_token"
)


func StartApplication() {

	atService := access_token.NewService(db.NeRepository())

}