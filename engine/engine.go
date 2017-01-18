package engine

import (
	common "github.com/hishboy/gocommons/lang"
	"fmt"
	"../jsonparser"
	"container/list"
)

type EngineContainer struct {
	queue 	*common.Queue
	cl	list.List
}

func NewEngine() *EngineContainer {
	ec 		:= &EngineContainer{}
	ec.queue 	= common.NewQueue()
	ec.cl		= list.List{}
	return ec
}


func (ec *EngineContainer) Start() {

	for {
		var msg ClientMessage
		if ec.queue.Len() > 0 {
			msg = ec.queue.Poll().(ClientMessage)
		}

		if msg.Message == "" {
			continue
		} else {
			out, processerror := ProcessMessage(msg)
			if processerror != nil {
				fmt.Println(msg.client.RemoteAddr(), "- " + processerror.Error())
				msg.client.PushMessage(jsonparser.Serialize(out))
			} else {
				msg.client.PushMessage(jsonparser.Serialize(out))
			}
		}
	}
}

func (ec *EngineContainer) PushMessage(msg interface{}) {

	ec.queue.Push(msg)

}

func (ec *EngineContainer) AddClient(client *ClientConnection){

	ec.cl.PushFront(client)

	for obj := ec.cl.Front(); obj != nil; obj = obj.Next(){
		//if obj.Value == client {
			obj.Value.(*ClientConnection).PushMessage("{\"Info\":\"Client - " + client.RemoteAddr() + " Joined\"}")
		//}
	}
}