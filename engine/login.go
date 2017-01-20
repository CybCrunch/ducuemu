package engine

// This doesn't really belong in the engine directory, it suffices for now

import (
	"fmt"
	parser "../jsonparser"
	"errors"
	"strconv"
	sanitize "../common"
)

func login(login []string, client *ClientConnection) (parser.JsonMessage, error) {

	if len(login) < 2 || len(login) > 2{
		return parser.Message("error",[]string{"Invalid login parameter count: " + strconv.Itoa(len(login))}),
			errors.New("Invalid login parameter count: " + strconv.Itoa(len(login)))
	}

	if len(login[0]) > 25 {
		return parser.Message("error", []string{"Username too long, max 25 characters"}),
			errors.New("Username too long, max 25 characters")
	}

	if len(login[1]) > 25 {
		return parser.Message("error", []string{"Invalid Password"}),
			errors.New("Invalid Password")
	}
	
	if sanitize.NonAscii(login[0]) {
		return parser.Message("error", []string{"Username contains invalid characters"}),
			errors.New("Username contains invalid characters")
	}

	fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Attempt - " + login[0])

	if client.user == ""{

		for u := client.ec.cl.Front(); u != nil; u = u.Next(){
			if u.Value.(*ClientConnection).user == login[0]{
				fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Failure (Already Logged In) - " + login[0])
				return parser.Message("error",[]string{"User is already logged in"}),
					errors.New("User is already logged in: " + login[0])
			}
		}

		dbh := client.ec.DB()

		if dbh.VerifyUser(login[0], login[1]){
			client.setUser(login[0])
			fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Success - " + login[0])
			client.ec.PushAll(parser.Message("chat", []string{login[0] + " has joined"}))
			return parser.Message("chat",[]string{"Welcome " + login[0] + "!"}), nil
		} else {
			fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Failure - " + login[0])
			return parser.Message("error",[]string{"Invalid User Credentials Supplied"}),
				errors.New("Invalid User Credentials - " + login[0])
		}



	} else {
		fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Failure (Already Logged In) - " + login[0])
		return parser.Message("error", []string{"You are already logged in!"}),
			errors.New("User already logged in: " + login[0])
	}
}
