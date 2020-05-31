package dvonn


type EdgeType int

const (
	UNKNOWN          EdgeType = iota
	UPPER_EDGE                // upper
	LEFTUPPER_EDGE            // left upper
	LEFTBOTTON_EDGE           // left button
	BOTTON_EDGE               // bottom
	RIGHTUPPER_EDGE           // right upper
	RIGHTBOTTON_EDGE          // right bottom
)
