package ships

import (
	el "../elements"
	p "../jsonparser"
)

type Ship struct {

	CU CoreUnit
	reactor	el.Reactor

}


func newShip() (*Ship) {

	sh := &Ship{}
	return sh

}

func (sh *Ship) ListElements() p.JsonMessage {

	return p.Message("error", []string{"Ship Type is Null"})

}

