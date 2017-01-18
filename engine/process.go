package engine

import (
	"fmt"
	"errors"
	"../jsonparser"
)


func ProcessMessage(message ClientMessage) (jsonparser.JsonMessage, error) {

	jm, err := jsonparser.Deserialize(message.Message)

	if err != nil {
		return jsonparser.JsonMessage{MessageType:"Error", Message:err.Error()},
			errors.New(err.Error())
	}

	if jm.MessageType == "event" {
		fmt.Println("Event Message Received")
		return jsonparser.JsonMessage{MessageType:"Event",
			Message:"Event Received Successfully"}, nil

	} else if jm.MessageType == "info" {
		fmt.Println("Info Message Received")
		return jsonparser.JsonMessage{MessageType:"Info",
			Message:"Info Received Successfully"}, nil
	} else {
		return jsonparser.JsonMessage{MessageType:"Error",
			Message:"Unknown Message Type: " + jm.MessageType },
			errors.New("Unknown Message Type: " + jm.MessageType)
	}

	return jm, nil
}
