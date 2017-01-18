package engine

import (
	common "github.com/hishboy/gocommons/lang"
	"fmt"
	"../jsonparser"
)

type EngineContainer struct {
	queue *common.Queue
}

func NewEngine() *EngineContainer {
	ec := &EngineContainer{}
	ec.queue = common.NewQueue()
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