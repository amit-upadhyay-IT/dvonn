package dvonn

import "log"

type MatchResult struct {
	WinningColor WinnerColor
	WinnerScore  int
	LoserScore   int
}

type WinnerColor string

const (
	WINNER_UNKNOWN WinnerColor = "UNKNOWN"
	WINNER_WHITE WinnerColor = "WHITE"
	WINNER_BLACK WinnerColor = "BLACK"
	WINNER_DRAW  WinnerColor = "DRAW"
)

func GetWinnerColorFromPlayerColor(color ChipColor) WinnerColor {
	if color == WHITE {
		return WINNER_WHITE
	} else if color == BLACK {
		return WINNER_BLACK
	}
	log.Fatal("player can not have color other than WHITE or BLACK")
	return WINNER_UNKNOWN
}

func (mr *MatchResult) GetWinnerColor() WinnerColor {
	return mr.WinningColor
}

//func (mr *MatchResult) GetWinnerScore() int {
//	return mr.WinnerScore
//}
//
//func (mr *MatchResult) GetLoserScore() int {
//	return mr.LoserScore
//}
