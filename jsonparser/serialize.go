package jsonparser

import "encoding/json"

func Serialize(jm JsonMessage) (string, error) {
	if message, err := json.Marshal(jm); err != nil{
		return "", err
	} else {
		return string(message), nil
	}
}
