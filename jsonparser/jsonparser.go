package jsonparser


func Message(mt string, m []string) JsonMessage {

	return JsonMessage{MessageType:mt, Message:m}

}