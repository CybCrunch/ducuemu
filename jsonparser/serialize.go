package jsonparser

import "encoding/json"

func Serialize(jm JsonMessage) string {

	message, _ := json.Marshal(jm)

	return string(message)
}
