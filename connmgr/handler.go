package connmgr

import (
	"bufio"
	"fmt"
	"io"
	"../engine"
)

func handleConnection(client *engine.ClientConnection) {

	ch  := make(chan string)
	eCh := make(chan error)

	go func(ch chan string, eCh chan error) {
		for {
			message, readerror := bufio.NewReader(client.Conn()).ReadString('\n')
			if readerror != nil {
				eCh<- readerror
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
			case err := <- eCh:
				if err != io.EOF{
					fmt.Println(client.Conn().RemoteAddr(), "- Error reading: ", err.Error())
				}
				client.Close()
				return
			default:
				break
		}

		if msg := client.PopMessage(); msg != nil {
			response := msg.(string)
			if _, err := client.Conn().Write([]byte(response + "\n")); err != nil {
				fmt.Println(client.RemoteAddr(), "- Error Writing: ", err.Error())
				client.Close()
				return
			}
		}
	}

}
