package dvonn

type GamePhase int

const (
	UNKNOWN_PHASE GamePhase = iota
	PLACEMENT_PHASE
	MOVEMENT_PHASE
)
