package db

import (
	"fmt"
	"github.com/yamamushi/ducuemu/common"
)

func (dbh *DBHandler) VerifyUser(userid string, password string) bool {

	rowpass, err := dbh.Query("Select password from users where username = '"+userid+"'")

	if err != nil {
		fmt.Println("Error Querying User - " + userid +": " + err.Error())
		return false
	}
	defer rowpass.Close()

	password = common.HashString(password)

	for rowpass.Next() {
		var result string
		err := rowpass.Scan(&result)
		if err != nil {
			fmt.Println("Error Querying User: " + err.Error())
			return false
		}

		if result == password {
			return true
		}
	}

	// We should never get to here
	return false
}
