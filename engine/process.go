package engine

import (
	"errors"
	parser "../jsonparser"
)


func ProcessMessage(cm ClientMessage) (parser.JsonMessage, error) {

	jm, err := parser.Deserialize(cm.Message)
	if err != nil {
		return parser.Message("error", []string{err.Error()}),
			errors.New(err.Error())
	}

	if jm.MessageType == "logout"{
		cm.client.PushMessage("Goodbye!")
		cm.client.PushMessage("close") // Tells our handler to close the connection accordingly
		return parser.JsonMessage{}, nil
	}

	if cm.client.user != "" {
		if jm.MessageType == "command" {
			return cm.client.ec.Command(cm.client, jm)
		} else if jm.MessageType == "login" {
			return parser.Message("error", []string{"You are already logged in!"}),
				errors.New("User already logged in: " + jm.MessageType)
		} else if jm.MessageType == "register" {
			return parser.Message("error", []string{"You are currently logged in!"}),
				errors.New("User already logged in and (presumably) registered: " + jm.MessageType)
		} else {
			return parser.Message("error", []string{"Unknown Message Type: " + jm.MessageType}),
				errors.New("Unknown Message Type: " + jm.MessageType)
		}
	} else if jm.MessageType == "login" {
		return login(jm.Message, cm.client)
	}  else if jm.MessageType == "register" {
		return register(jm.Message, cm.client)
	} else {
		return parser.Message("error", []string{"You must register and login before continuing"}), nil
	}
	return parser.JsonMessage{}, nil
}
