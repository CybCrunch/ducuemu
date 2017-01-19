package engine

import (
	"fmt"
	parser "../jsonparser"
	"errors"
	"strconv"
)

func register(credentials []string, client *ClientConnection) (parser.JsonMessage, error) {

	if len(credentials) < 2 || len(credentials) > 2{
		return parser.JsonMessage{MessageType:"error",
			Message:[]string{"Invalid registration parameter count: " + strconv.Itoa(len(credentials))}},
			errors.New("Invalid registration parameter count: " + strconv.Itoa(len(credentials)))
	}

	fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Attempt - " + credentials[0])

	if client.user == ""{

		for u := client.ec.cl.Front(); u != nil; u = u.Next(){
			if u.Value.(*ClientConnection).user == credentials[0]{
				fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Failure (Already Logged In) - " + credentials[0])
				return parser.JsonMessage{MessageType:"error",
					Message:[]string{"User is already logged in"}},
					errors.New("User is already logged in: " + credentials[0])
			}
		}

		dbh := client.ec.DB()

		usermatch, err := dbh.Query("SELECT COUNT(*) as count FROM users where username ='"+credentials[0]+"'")

		if err != nil {

			fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Failure - " + credentials[0] + " - " + err.Error())
			return parser.JsonMessage{MessageType:"error",
				Message:[]string{"User Registration Failure"}}, err

		} else if count, _ := dbh.CheckCount(usermatch); count > 0 {

			fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Failure (Already Exists) - " + credentials[0])
			return parser.JsonMessage{MessageType:"error",
				Message:[]string{"Username already taken"}},
				errors.New("Username is already taken: " + credentials[0])
		}


		if err := dbh.RegisterUser(credentials[0], credentials[1]); err != nil {
			fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Failure - " + credentials[0] + " - " + err.Error())
			return parser.JsonMessage{MessageType:"info",
				Message:[]string{"Registration Success for " + credentials[0] + "!"}}, err

		} else {

			fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Success - " + credentials[0])
			return parser.JsonMessage{MessageType:"info",
				Message:[]string{"Registration Success for " + credentials[0] + "!"}}, nil
		}

	} else {
		fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Failure (Already Logged In) - " + credentials[0])
		return parser.JsonMessage{MessageType:"error",
			Message:[]string{"You are currently logged in!"}}, errors.New("User already logged in: " + credentials[0])
	}
}
