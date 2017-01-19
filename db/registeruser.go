package db

import (
	"../common"
)


func (dbh *DBHandler) RegisterUser(userid string, password string){

	sha := common.HashString(password)

	dbh.Query(" INSERT INTO `users` (`id`,`username`,`password`,`email`) VALUES"+
	"('','"+userid+"','"+sha+"','');")

}
