package connmgr

import (
	"fmt"
	"net"
	"os"
	"../engine"
)

func Start(ci ConnInfo, ec *engine.EngineContainer) {

	l, err := net.Listen(ci.CONN_TYPE, ci.CONN_HOST+":"+ci.CONN_PORT)
	if err != nil {
		fmt.Println("Error Establishing Listen Socket:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + ci.CONN_HOST + ":" + ci.CONN_PORT)

	for {

		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error Accepting: ", err.Error())
			os.Exit(1)
		}

		conn := engine.NewClient(c, ec)

		go handleConnection(conn)
	}
}
