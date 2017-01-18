package jsonparser

import (
	"encoding/json"
	"errors"
	"strings"
)

func Deserialize(message string) (JsonMessage, error) {

	res := JsonMessage{}

	if err := json.Unmarshal([]byte(message), &res); err != nil {
		return JsonMessage{}, errors.New("Error Deserializing Message: " + strings.TrimSpace(message))
	}

	return res, nil

}
