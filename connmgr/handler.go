package connmgr

import (
	"bufio"
	"fmt"
	"io"
	"../engine"
)

func handleConnection(client *engine.ClientConnection) {

	// Loop until connection is closed
	for {
		message, readerror := bufio.NewReader(client.Conn()).ReadString('\n')
		if readerror != nil {
			if readerror == io.EOF {
				fmt.Println(client.Conn().RemoteAddr(), "- EOF Detected")
			} else {
				fmt.Println(client.Conn().RemoteAddr(), "- Error reading: ", readerror.Error())
			}
			CloseConnection(client.Conn())
			return
		}

		// Process message received
		client.Process(message)

		var response string

		for {
			if err := client.PopMessage(); err != nil {
				response = err.(string)
				break
			}
		}

		// Send a response back to person contacting us.
		if _, err := client.Conn().Write([]byte(response + "\n")); err != nil {
			fmt.Println(client.RemoteAddr(), "- Error Writing: ", err.Error())
			CloseConnection(client.Conn())
			return
		}
	}

}
