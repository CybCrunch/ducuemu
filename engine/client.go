package engine

import (
	"net"
	common "github.com/hishboy/gocommons/lang"
	"fmt"
	sh "../ships"
	parser "../jsonparser"
)


type ClientConnection struct {

	conn net.Conn
	queue *common.Queue
	ec *EngineContainer
	user string
	ship interface{}

}

type ClientMessage struct {

	client *ClientConnection
	Message	string

}


func NewClient(conn net.Conn, ec *EngineContainer) *ClientConnection {

	client := &ClientConnection{conn, common.NewQueue(), ec, "", sh.NewShip()}
	ec.AddClient(client)
	return client

}


func (client *ClientConnection) RemoteAddr() string {

	return client.conn.RemoteAddr().String()

}

func (client *ClientConnection) PushMessage(msg interface{}) {
	client.queue.Push(msg)
}

func (client *ClientConnection) PopMessage() interface{} {

	return client.queue.Poll()

}

func (client *ClientConnection) Conn() net.Conn {

	return client.conn

}

func (client *ClientConnection) Process(msg string){
	out := ClientMessage{client, msg}
	client.ec.PushMessage(out)
}

// A handler for cleaning up after closing a socket
func (client *ClientConnection) Close() {
	fmt.Println(client.RemoteAddr(), "- Connection Closed")
	client.ec.RemoveClient(client)
	client.Conn().Close()
}

func (client *ClientConnection) setUser(userid string){
	client.user = userid
}

func (client *ClientConnection) InitClient(username string) error {

	client.user = username
	client.LoadShip()

	client.ec.PushAll(parser.Message("chat", []string{username + " has joined"}))
	return nil
}

func (client *ClientConnection) CleanupClient() {

	client.SaveShip()
}

func (client *ClientConnection) LoadShip() {

	client.ship = sh.NewCruiser()

}

func (client *ClientConnection) SaveShip() {


}