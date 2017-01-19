package engine

import (
	"fmt"
	parser "../jsonparser"
)
func Login(login string, client *ClientConnection) parser.JsonMessage{

	fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Attempt - " + login)

	if client.user == ""{
		for u := client.ec.cl.Front(); u != nil; u = u.Next(){
			if u.Value.(*ClientConnection).user == login{
				return parser.JsonMessage{MessageType:"error",
					Message:[]string{"Username has been taken, please select another"}}
			}
		}
		client.setUser(login)
		fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Success - " + login)
		client.ec.PushAll(parser.JsonMessage{MessageType:"chat", Message:[]string{login + " has joined"}})
		return parser.JsonMessage{MessageType:"chat",
			Message:[]string{"Welcome " + login + "!"}}

	} else {
		return parser.JsonMessage{MessageType:"error",
			Message:[]string{"You are already logged in!"}}
	}
}
