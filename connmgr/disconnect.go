package connmgr

import (
	"fmt"
	"net"
)

// A handler for cleaning up after closing a socket
func CloseConnection(conn net.Conn) {

	fmt.Println(conn.RemoteAddr(), "- Connection Closed")
	conn.Close()

}
