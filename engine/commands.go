package engine

import (
	parser "../jsonparser"
	"fmt"
	"errors"
	"strconv"
)


func (ec *EngineContainer) Command(client *ClientConnection, jm parser.JsonMessage) (parser.JsonMessage, error) {

	jml := len(jm.Message)
	message := jm.Message
	command := jm.MessageType

	if command == "set" {
		return setCommand()

	} else if command == "chat" {
		return chatCommand(client, jm)
	} else if command == "info" {
		return infoCommand(client, jm)
	}

	return parser.Message("error",
			[]string{"Invalid Command: " + message[0]}),
			errors.New(client.RemoteAddr() + " - Invalid command: " + message[0])

}

func voidCommand(client *ClientConnection, jm parser.JsonMessage) (parser.JsonMessage, error) {
	return parser.Message("", []string{}), nil
}

func setCommand(client *ClientConnection, jm parser.JsonMessage) (parser.JsonMessage, error) {

	return parser.Message("", []string{}), nil
}

func infoCommand(client *ClientConnection, jm parser.JsonMessage) (parser.JsonMessage, error) {
	return parser.Message("", []string{}), nil
}

func chatCommand(client *ClientConnection, jm parser.JsonMessage) (parser.JsonMessage, error) {

	jml := len(jm.Message)
	message := jm.Message

	if jml != 2 {
		return parser.Message("error",
			[]string{"Invalid Number of parameters for command [chat] expected 2 got " + strconv.Itoa(jml)}),
			errors.New(client.RemoteAddr() + " - Invalid parameter count")
	}

	fmt.Println("[Chat Log]:[" + client.RemoteAddr() + "] - " + message[1])
	client.ec.PushAllFrom(client, parser.Message("chat", []string{client.user, message[1]}))
	return parser.Message("", []string{}), nil
}