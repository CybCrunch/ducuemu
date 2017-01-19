package db

import (
	"../common"
	"fmt"
)


func (dbh *DBHandler) RegisterUser(userid string, password string) (error) {

	sha := common.HashString(password)


	_, err := dbh.Query(" INSERT INTO `users` (`id`,`username`,`password`,`email`) VALUES"+
	"('','"+userid+"','"+sha+"','');")

	if err != nil {
		return err
	} else {
		fmt.Println("Registered User: " + userid + " || " + password + " || " + sha)
		return nil
	}
}
