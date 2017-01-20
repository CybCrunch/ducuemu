package engine

import (
	"fmt"
	parser "../jsonparser"
	"container/list"
	"../config"
	"../db"
)

type EngineContainer struct {
	msgchan	chan(ClientMessage)
	cl	list.List
	cfg	*config.DucuemuConfig
	dbh	*db.DBHandler
}

func NewEngine(ducfg *config.DucuemuConfig, dbh *db.DBHandler) *EngineContainer {
	ec 		:= &EngineContainer{}
	ec.msgchan	= make(chan ClientMessage)
	ec.cl		= list.List{}
	ec.cfg		= ducfg
	ec.dbh		= dbh
	return ec
}


func (ec *EngineContainer) Start() {
	fmt.Println("Game Engine Started")
	for {
		select {

		case msg := <- ec.msgchan:
			processed, processerror := ProcessMessage(msg)
			if processerror != nil {
				fmt.Println(msg.client.RemoteAddr(), "- " + processerror.Error())
				out, _ := parser.Serialize(processed)
				msg.client.PushMessage(out)
			} else {
				out, _ := parser.Serialize(processed)
				if processed.MessageType == ""{
					break
				}
				msg.client.PushMessage(out)
			}
		default:
			break
		}
	}
	fmt.Println("Game Engine Shutdown")
}

func (ec *EngineContainer) PushMessage(msg interface{}) {
	ec.msgchan <- msg.(ClientMessage)
}

func (ec *EngineContainer) AddClient(client *ClientConnection){
	ec.cl.PushFront(client)
	fmt.Println("[user]: " + client.RemoteAddr() + " Connected")
}

func (ec *EngineContainer) RemoveClient(client *ClientConnection){
	for e := ec.cl.Front(); e != nil; e = e.Next() {
		if e.Value == client {
			ec.cl.Remove(e)
		}
	}
	ec.PushAll(parser.Message("chat", []string{client.user + " has disconnected."}))
}

func (ec *EngineContainer) PushAll(msg interface{}) error {
	if out, err := parser.Serialize(msg.(parser.JsonMessage)); err != nil{
		return err
	} else{
		for obj := ec.cl.Front(); obj != nil; obj = obj.Next(){
			if obj.Value.(*ClientConnection).user != ""{
				obj.Value.(*ClientConnection).PushMessage(out)
			}
		}
		return nil
	}
}

func (ec *EngineContainer) DB() (*db.DBHandler) {
	return ec.dbh
}