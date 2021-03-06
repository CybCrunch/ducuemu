package engine

// This doesn't really belong in the engine directory, it suffices for now

import (
	"fmt"
	parser "../jsonparser"
	"errors"
	"strconv"
	sanitize "../common"
)

func register(register []string, client *ClientConnection) (parser.JsonMessage, error) {

	if len(register) < 2 || len(register) > 2{
		return parser.Message("error", []string{"Invalid registration parameter count: " + strconv.Itoa(len(register))}),
			errors.New("Invalid registration parameter count: " + strconv.Itoa(len(register)))
	}

	if len(register[0]) > 25 {
		return parser.Message("error", []string{"Username too long, max 25 characters"}),
			errors.New("Username too long, max 25 characters")
	}

	if len(register[1]) > 25 {
		return parser.Message("error", []string{"Password too long, max 25 characters"}),
			errors.New("Password too long, max 25 characters")
	}

	if sanitize.NonAscii(register[0]) {
		return parser.Message("error", []string{"Username contains invalid characters"}),
			errors.New("Username contains invalid characters")
	}



	fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Attempt - " + register[0])

	if client.user == ""{

		for u := client.ec.cl.Front(); u != nil; u = u.Next(){
			if u.Value.(*ClientConnection).user == register[0]{
				fmt.Println("[user]: " + client.RemoteAddr() + " : User Login Failure (Already Logged In) - " + register[0])
				return parser.Message("error", []string{"User is already logged in"}),
					errors.New("User is already logged in: " + register[0])
			}
		}

		dbh := client.ec.DB()

		usermatch, err := dbh.Query("SELECT COUNT(*) as count FROM users where username ='"+register[0]+"'")

		if err != nil {

			fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Failure - " + register[0] + " - " + err.Error())
			return parser.Message("error", []string{"User Registration Failure"}), err

		} else if count, _ := dbh.CheckCount(usermatch); count > 0 {
			fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Failure (Already Exists) - " + register[0])
			return parser.Message("error", []string{"Username already taken"}),
				errors.New("Username is already taken: " + register[0])
		}

		if err := dbh.RegisterUser(register[0], register[1]); err != nil {
			fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Failure - " + register[0] + " - " + err.Error())
			return parser.Message("info", []string{"Registration Success for " + register[0] + "!"}), err

		} else {

			fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Success - " + register[0])
			return parser.Message("info", []string{"Registration Success for " + register[0] + "!"}), nil
		}

	} else {
		fmt.Println("[user]: " + client.RemoteAddr() + " : User Registration Failure (Already Logged In) - " + register[0])
		return parser.Message("error", []string{"You are currently logged in!"}), errors.New("User already logged in: " + register[0])
	}
}
