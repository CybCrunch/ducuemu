package engine

import (
	"fmt"
	parser "../jsonparser"
	"errors"
	"strconv"
)

func login(login []string, client *ClientConnection) (parser.JsonMessage, error) {

	if len(login) < 2 || len(login) > 2{
		return parser.JsonMessage{MessageType:"error",
			Message:[]string{"Invalid login parameter count: " + strconv.Itoa(len(login))}},
			errors.New("Invalid login parameter count: " + strconv.Itoa(len(login)))
	}

	fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Attempt - " + login[0])

	if client.user == ""{

		for u := client.ec.cl.Front(); u != nil; u = u.Next(){
			if u.Value.(*ClientConnection).user == login[0]{
				fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Failure (Already Logged In) - " + login[0])
				return parser.JsonMessage{MessageType:"error",
					Message:[]string{"User is already logged in"}},
					errors.New("User is already logged in: " + login[0])
			}
		}

		dbh := client.ec.DB()

		if dbh.VerifyUser(login[0], login[1]){
			client.setUser(login[0])
			fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Success - " + login[0])
			client.ec.PushAll(parser.JsonMessage{MessageType:"chat", Message:[]string{login[0] + " has joined"}})
			return parser.JsonMessage{MessageType:"chat",
				Message:[]string{"Welcome " + login[0] + "!"}}, nil
		} else {
			fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Failure - " + login[0])
			return parser.JsonMessage{MessageType:"error",
				Message:[]string{"Invalid User Credentials Supplied"}}, errors.New("Invalid User Credentials - " + login[0])
		}



	} else {
		fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Failure (Already Logged In) - " + login[0])
		return parser.JsonMessage{MessageType:"error",
			Message:[]string{"You are already logged in!"}}, errors.New("User already logged in: " + login[0])
	}
}
