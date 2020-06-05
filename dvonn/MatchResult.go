package dvonn

import "log"

type MatchResult string

const (
	RESULT_WHITE MatchResult = "WHITE"
	RESULT_BLACK MatchResult = "BLACK"
	RESULT_DRAW MatchResult = "DRAW"
)


func GetMatchResultFromPlayerColor(color ChipColor) MatchResult {
	var res MatchResult
	if color == WHITE {
		res = RESULT_WHITE
	} else if color == BLACK {
		res = RESULT_BLACK
	}
	log.Fatal("player can not have color other than WHITE or BLACK")
	return res
}
