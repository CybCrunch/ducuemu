package main

import (
	"../../../ducuemu/connmgr"
	"../../engine"
)

func main() {

	ec := engine.NewEngine()
	go ec.Start()

	ci := connmgr.ConnInfo{CONN_HOST: "localhost", CONN_PORT: "8001", CONN_TYPE: "tcp"}
	connmgr.Start(ci, ec)

}
