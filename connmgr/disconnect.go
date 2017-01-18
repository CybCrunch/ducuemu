package connmgr

import (
	"fmt"
	"../engine"

)

// A handler for cleaning up after closing a socket
func CloseConnection(client *engine.ClientConnection) {

	fmt.Println(client.RemoteAddr(), "- Connection Closed")
	client.Conn().Close()

}
