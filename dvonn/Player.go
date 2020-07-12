package dvonn

import "log"

type Player struct {
	id          string
	name        string
	playerColor ChipColor // this can only be WHITE or BLACK
}

func GetPlayer(name, id string, playerColor ChipColor) Player {
	if playerColor == RED {
		log.Fatal("wrong color is passed while creating player")
	}
	return Player{id, name, playerColor}
}

func (p Player) GetId() string {
	return p.id
}

func (p Player) GetName() string {
	return p.name
}

func (p Player) GetPlayerColor() ChipColor {
	return p.playerColor
}
