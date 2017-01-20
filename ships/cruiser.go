package ships

import (
	el "../elements"
	p "../jsonparser"
)


type Cruiser struct {

	Ship

	ft el.Thruster	// Forward Thruster (located on the front of the ship)
	rt el.Thruster  // Rear Thruster (located on the rear of the ship)
	ldt el.Thruster // Left Directional Thruster
	rdt el.Thruster // Right Directional Thruster
	udt el.Thruster // Upper Directional Thruster
	bdt el.Thruster // Bottom Directional Thruster

	ll  el.Launcher // Left Launcher
	rl  el.Launcher // Right Launcher
	scanner el.Scanner

}

func NewCruiser() *Cruiser {
	cruiser := &Cruiser{}
	return cruiser
}


func (sh *Cruiser) ListElements() p.JsonMessage {

	return p.Message("error", []string{
		"reactor", "scanner",
		"left_launcher", "right_launcher",
		"forward_thruster", "rear_thruster",
		"left_d_thruster", "right_d_thruster",
		"upper_d_thruster", "lower_d_thruster"})

}

