package engine

import (
	"fmt"
	"errors"
	parser "../jsonparser"
)


func ProcessMessage(cm ClientMessage) (parser.JsonMessage, error) {

	jm, err := parser.Deserialize(cm.Message)

	if err != nil {
		return parser.JsonMessage{MessageType:"error", Message:[]string{err.Error()}},
			errors.New(err.Error())
	}

	if cm.client.user != "" {

		if jm.MessageType == "event" {

			fmt.Println("Event Message Received")
			return parser.JsonMessage{MessageType:"event",
				Message:[]string{"Event Received Successfully"}}, nil

		} else if jm.MessageType == "info" {

			fmt.Println("Info Message Received")
			return parser.JsonMessage{MessageType:"info",
				Message:[]string{"Info Received Successfully"}}, nil

		} else if jm.MessageType == "chat" {

			fmt.Println("[Chat Log]:[" + cm.client.RemoteAddr() + "] - " + jm.Message[1])
			cm.client.ec.PushAll(parser.JsonMessage{MessageType:"chat",
				Message:[]string{cm.client.user, jm.Message[1] }})
			return parser.JsonMessage{}, nil

		} else if jm.MessageType == "login" {
			return parser.JsonMessage{MessageType:"error",
				Message:[]string{"You are already logged in!"}}, nil
		} else {

			return parser.JsonMessage{MessageType:"error",
				Message:[]string{"Unknown Message Type: " + jm.MessageType} },
				errors.New("Unknown Message Type: " + jm.MessageType)
		}

	} else if jm.MessageType == "login" {
		return login(jm.Message, cm.client)
	}  else if jm.MessageType == "register" {
		return register(jm.Message, cm.client)
	} else {

		return parser.JsonMessage{MessageType:"error",
			Message:[]string{"You must register and login before continuing"}}, nil
	}


	return parser.JsonMessage{}, nil
}
