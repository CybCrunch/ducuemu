package engine

import (
	"fmt"
	"../jsonparser"
	"container/list"
)

type EngineContainer struct {
	msgchan	chan(ClientMessage)
	cl	list.List
}

func NewEngine() *EngineContainer {
	ec 		:= &EngineContainer{}
	ec.msgchan	= make(chan ClientMessage)
	ec.cl		= list.List{}
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
				out, _ := jsonparser.Serialize(processed)
				msg.client.PushMessage(out)
			} else {
				out, _ := jsonparser.Serialize(processed)
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
	msg := jsonparser.JsonMessage{MessageType:"info", Message:[]string{"Client Joined - " + client.RemoteAddr()}}
	ec.PushAll(msg)
}

func (ec *EngineContainer) RemoveClient(client *ClientConnection){
	for e := ec.cl.Front(); e != nil; e = e.Next() {
		if e.Value == client {
			ec.cl.Remove(e)
		}
	}
}

func (ec *EngineContainer) PushAll(msg interface{}) error {
	if out, err := jsonparser.Serialize(msg.(jsonparser.JsonMessage)); err != nil{
		return err
	} else{
		for obj := ec.cl.Front(); obj != nil; obj = obj.Next(){
			obj.Value.(*ClientConnection).PushMessage(out)
		}
		return nil
	}
}