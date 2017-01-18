package connmgr

import (
	"bufio"
	"fmt"
	"io"
	"../engine"
)

func handleConnection(client *engine.ClientConnection) {

	// Loop until connection is closed

	ch  := make(chan string)
	eCh := make(chan error)

	go func(ch chan string, eCh chan error) {
		for {
			message, readerror := bufio.NewReader(client.Conn()).ReadString('\n')
			if readerror != nil {
				if readerror == io.EOF {
					fmt.Println(client.Conn().RemoteAddr(), "- EOF Detected")
					eCh<- readerror
				} else {
					fmt.Println(client.Conn().RemoteAddr(), "- Error reading: ", readerror.Error())
					eCh<- readerror
				}
				CloseConnection(client)
				return
			}
			ch<- message
		}
	}(ch, eCh)

	for {
		// Process message received
		select {
			case message := <-ch:
				client.Process(message)
			default:
				break
		}

		var response string
		if err := client.PopMessage(); err != nil {
			response = err.(string)
			// Send a response back to person contacting us.
			if _, err := client.Conn().Write([]byte(response + "\n")); err != nil {
				fmt.Println(client.RemoteAddr(), "- Error Writing: ", err.Error())
				CloseConnection(client)
				return
			}
		}
		//break
	}

}
