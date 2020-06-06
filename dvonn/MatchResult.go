package dvonn

import "log"

type MatchResult struct {
	winningColor WinnerColor
	winnerScore int
	loserScore int
}

type WinnerColor string

const (
	WINNER_WHITE WinnerColor = "WHITE"
	WINNER_BLACK WinnerColor = "BLACK"
	WINNER_DRAW  WinnerColor = "DRAW"
)

func GetWinnerColorFromPlayerColor(color ChipColor) WinnerColor {
	var res WinnerColor
	if color == WHITE {
		res = WINNER_WHITE
	} else if color == BLACK {
		res = WINNER_BLACK
	}
	log.Fatal("player can not have color other than WHITE or BLACK")
	return res
}

func (mr *MatchResult) GetWinnerColor() WinnerColor {
	return mr.winningColor
}

func (mr *MatchResult) GetWinnerScore() int {
	return mr.winnerScore
}

func (mr *MatchResult) GetLoserScore() int {
	return mr.loserScore
}
