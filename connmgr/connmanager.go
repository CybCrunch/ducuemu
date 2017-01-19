package connmgr

import (
	"fmt"
	"net"
	"os"
	"../engine"
	"../config"
)

func Start(cfg *config.DucuemuConfig, ec *engine.EngineContainer) {


	l, err := net.Listen("tcp", cfg.Host())
	if err != nil {
		fmt.Println("Error Establishing Listen Socket:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + cfg.Host())

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
