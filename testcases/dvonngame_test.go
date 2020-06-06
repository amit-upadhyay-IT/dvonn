package testcases

import (
	"../dvonn"
	"fmt"
	"strings"
	"testing"
)

func TestGameSample_01(t *testing.T) {
	/*
	         5
	      4   \ __
	   3   \ __/w \__
	    \ __/b \__/b \__
	  2  /b \__/b \__/b \__
	   \ \__/b \__/w \__/b \__
	  1  /w \__/b \__/w \__/w \__
	   \ \__/w \__/b \__/w \__/w \__
	     /w \__/b \__/w \__/b \__/b \__
	     \__/D \__/w \__/D \__/b \__/D \__
	      | \__/w \__/w \__/b \__/w \__/b \
	      A  | \__/b \__/b \__/w \__/b \__/
	         B  | \__/b \__/b \__/w \__/b \
	            C  | \__/w \__/b \__/w \__/
	               D  | \__/w \__/w \__/w \
	                  E  | \__/w \__/w \__/
	                     F  | \__/b \__/|
	                        G  | \__/|  K
	                           H  |  J
	                              I
	*/

	players := make([]dvonn.Player, 0)
	player1 := dvonn.GetPlayer("pushpender", "amit123", dvonn.WHITE)
	player2 := dvonn.GetPlayer("arman", "amit123", dvonn.BLACK)
	players = append(players, player1)
	players = append(players, player2)
	dvonnGame := dvonn.GetDvonnGame(players, player1)
	whitePlacementMoves := []string{
		"f1", "e4", "E5", "D5", "E2", "K3", "J2", "I3", "D2", "D3", "E3", "H5", "G5", "D1", "I5", "A3", "I1", "C2", "B1", "H1", "F2", "D4", "K4", "F4", "B3",
	}
	blackPlacementMoves := []string{
		"G3", "K5", "C1", "J4", "E1", "H4", "J5", "A1", "G4", "J3", "F5", "C3", "B2", "G2", "B4", "C5", "G1", "A2", "I4", "H2", "I2", "C4", "F3", "H3",
	}
	whiteMovementMoves := []string{
		"K3 J2", "G5 F5", "J2 H2", "I1 I2", "H5 G4", "K4 J3", "D2 E3", "D5 C5", "E2 F2", "G4 G2", "F4 F5", "I2 G2", "F5 F2", "E5 E4", "I3 H2", "E4 C2", "C2 G2",
	}
	blackMovementMoves := []string{
		"C1 D1", "G1 H1", "H1 H3", "J5 K5", "D1 D3", "B4 B3", "B3 D3", "J4 J3", "B2 C2", "C3 D4", "I4 H4", "E1 F2", "A1 B1", "H4 F2", "F3 E3", "E3 H3",
	}
	currentPlayer := player1
	whitePlacementCounter := 0
	blackPlacementCounter := 0
	whiteMovementCounter := 0
	blackMovementCounter := 0
	for {
		moveIds := make([]string, 0)
		if dvonnGame.GetGamePhase() == dvonn.PLACEMENT_PHASE {
			if currentPlayer.GetPlayerColor() == dvonn.WHITE {
				moveIds = append(moveIds, strings.ToLower(whitePlacementMoves[whitePlacementCounter]))
				whitePlacementCounter += 1
			} else {
				moveIds = append(moveIds, strings.ToLower(blackPlacementMoves[blackPlacementCounter]))
				blackPlacementCounter += 1
			}
		} else {
			if currentPlayer.GetPlayerColor() == dvonn.WHITE {
				moveIds = append(moveIds, strings.Split(strings.ToLower(whiteMovementMoves[whiteMovementCounter]), " ")...)
				whiteMovementCounter += 1
			} else {
				moveIds = append(moveIds, strings.Split(strings.ToLower(blackMovementMoves[blackMovementCounter]), " ")...)
				blackMovementCounter += 1
			}
		}
		moveRes := dvonnGame.Move(currentPlayer, moveIds...)
		if moveRes.IsGameOver() {
			break
		}
		if !moveRes.IsActionSuccess() {
			fmt.Println(moveRes.GetErrorCode())
			fmt.Println(moveRes.GetErrorMessage())
			if currentPlayer.GetPlayerColor() == dvonn.WHITE {
				whiteMovementCounter -= 1
			} else {
				blackMovementCounter -= 1
			}
		}
		currentPlayer = moveRes.GetNextPlayer()
	}
	winner, err := dvonnGame.GetGameWinner()
	if winner == nil || err != nil {
		// game should have ended
		t.Error("game should have ended by now")
	}
	if winner != nil && winner.GetWinnerColor() != dvonn.WINNER_BLACK {
		t.Error("black should have won")
	}

	if winner != nil && winner.GetLoserScore() != 13 {
		t.Error("loser score should be 13")
	}

	if winner != nil && winner.GetWinnerScore() != 14 {
		t.Error("winner score should be 14")
	}
}
