package ships

import (
	el "../elements"
	p "../jsonparser"
)

type Ship struct {

	Name string
	CU CoreUnit
	reactor	el.Reactor

}


func NewShip() (*Ship) {

	sh := &Ship{}
	return sh

}

func (sh *Ship) ListElements() p.JsonMessage {

	return p.Message("error", []string{"Ship Type is Null"})

}

func (sh *Ship) SetName(s string){
	sh.Name = s
}