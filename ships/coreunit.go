package ships


// We are using a CoreUnit here as a central place of storing information about our ship
// It's not really representative of what a Core Unit will actually do in Dual Universe

type CoreUnit struct {

	posX float64
	posY float64

	rotX	float64
	rotY	float64
	rotZ	float64

	speed	float64

	power	float64

}

