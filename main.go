package main

import (
	"./connmgr"
	"./engine"
	"./config"
	"./db"
)

func main() {

	config := config.NewConfig("")
	config.Read()

	dbh := db.NewDBHandler(config)
	dbh.Setup()
	//dbh.RegisterUser("testuser", "password")
	//dbh.VerifyUser("testuser", "W6ph5Mm5Pz8GgiULbPgzG37mj9g=")

	ec := engine.NewEngine(config, dbh)
	go ec.Start()

	connmgr.Start(config, ec)

}
