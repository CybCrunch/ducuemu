package engine

import (
	"fmt"
	"errors"
	"../jsonparser"
)


func ProcessMessage(cm ClientMessage) (jsonparser.JsonMessage, error) {

	jm, err := jsonparser.Deserialize(cm.Message)

	if err != nil {
		return jsonparser.JsonMessage{MessageType:"error", Message:[]string{err.Error()}},
			errors.New(err.Error())
	}

	if cm.client.user != "" {

		if jm.MessageType == "event" {

			fmt.Println("Event Message Received")
			return jsonparser.JsonMessage{MessageType:"event",
				Message:[]string{"Event Received Successfully"}}, nil

		} else if jm.MessageType == "info" {

			fmt.Println("Info Message Received")
			return jsonparser.JsonMessage{MessageType:"info",
				Message:[]string{"Info Received Successfully"}}, nil

		} else if jm.MessageType == "chat" {

			fmt.Println("[Chat Log]:[" + cm.client.RemoteAddr() + "] - " + jm.Message[1])
			cm.client.ec.PushAll(jsonparser.JsonMessage{MessageType:"chat",
				Message:[]string{cm.client.user, jm.Message[1] }})
			return jsonparser.JsonMessage{}, nil

		} else if jm.MessageType == "login" {
			return jsonparser.JsonMessage{MessageType:"error",
				Message:[]string{"You are already logged in!"}}, nil
		} else {

			return jsonparser.JsonMessage{MessageType:"error",
				Message:[]string{"Unknown Message Type: " + jm.MessageType} },
				errors.New("Unknown Message Type: " + jm.MessageType)
		}

	} else if jm.MessageType == "login" {
		return Login(jm.Message[0], cm.client), nil
	} else {

		return jsonparser.JsonMessage{MessageType:"error",
			Message:[]string{"You must login before continuing"}}, nil
	}


	return jsonparser.JsonMessage{}, nil
}
