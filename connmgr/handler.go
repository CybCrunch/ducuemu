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
			select {
			default:
				message, readerror := bufio.NewReader(client.Conn()).ReadString('\n')
				if readerror != nil {
					eCh <- readerror
					return
				}
				ch <- message
			}
		}
	}(ch, eCh)

	go func(ch chan string, eCh chan error){
		// Process message received
		for {
			select {
			case message := <-ch:
				client.Process(message)
			case err := <-eCh:
				if err != io.EOF {
					fmt.Println(client.Conn().RemoteAddr(), "- Error reading: ", err.Error())
				}
				client.Close()
				return
			case out := <-client.Mch:
				if out == "close" {
					fmt.Println(client.RemoteAddr(), " - Close Session Requested" )
					client.Close()
					return
				}
				if _, err := client.Conn().Write([]byte(out + "\n")); err != nil {
					fmt.Println(client.RemoteAddr(), "- Error Writing: ", err.Error())
					client.Close()
					return
				}
			}
		}
	}(ch, eCh)

}
