package db

import "fmt"

func (dbh *DBHandler) CreateTables() error {
	_, err := dbh.Query("CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER NULL AUTO_INCREMENT DEFAULT NULL,"+
		"`username` MEDIUMTEXT NULL DEFAULT NULL, `password` MEDIUMTEXT NULL DEFAULT NULL, `email`"+
		" MEDIUMTEXT NULL DEFAULT NULL, PRIMARY KEY (`id`));")
	if err != nil {
		fmt.Println("Error Creating User Table: " + err.Error())
		return err
	}

	_, err = dbh.Query("alter table users add UNIQUE(username(255));")
	if err != nil {
		fmt.Println("Error Altering User Table: " + err.Error())
		return err
	}

	return nil
}

