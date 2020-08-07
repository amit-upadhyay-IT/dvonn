package testcases

import (
	"../dvonn"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

/*
  White:A        Black:B

  0.  F1
  1.            G3
  2.  E4
  3.            K5
  4.  E5
  5.            C1
  6.  D5
  7.            J4
  8.  E2
  9.            E1
 10.  K3
 11.            H4
 12.  J2
 13.            J5
 14.  I3
 15.            A1
 16.  D2
 17.            G4
 18.  D3
 19.            J3
 20.  E3
 21.            F5
 22.  H5
 23.            C3
 24.  G5
 25.            B2
 26.  D1
 27.            G2
 28.  I5
 29.            B4
 30.  A3
 31.            C5
 32.  I1
 33.            G1
 34.  C2
 35.            A2
 36.  B1
 37.            I4
 38.  H1
 39.            H2
 40.  F2
 41.            I2
 42.  D4
 43.            C4
 44.  K4
 45.            F3
 46.  F4
 47.            H3
 48.  B3
 49.  K3J2
 50.            C1D1
 51.  G5F5
 52.            G1H1
 53.  J2H2
 54.            H1H3
 55.  I1I2
 56.            J5K5
 57.  H5G4
 58.            D1D3
 59.  K4J3
 60.            B4B3
 61.  D2E3
 62.            B3D3
 63.  D5C5
 64.            J4J3
 65.  E2F2
 66.            B2C2
 67.  G4G2
 68.            C3D4
 69.  F4F5
 70.            I4H4
 71.  I2G2
 72.            E1F2
 73.  F5F2
 74.            A1B1
 75.  E5E4
 76.            H4F2
 77.  I3H2
 78.            F3E3
 79.  E4C2
 80.            E3H3
 81.  C2G2

Black wins.

Status:   LOSE  WIN
Game pts: 13  104
*/

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
	player1 := dvonn.GetPlayer("pushpender", "amit1234", dvonn.WHITE)
	player2 := dvonn.GetPlayer("arman", "1234amit", dvonn.BLACK)
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

	if winner != nil && winner.LoserScore != 13 {
		t.Error("loser score should be 13")
	}

	if winner != nil && winner.WinnerScore != 14 {
		t.Error("winner score should be 14")
	}
}

type GamePlayStore struct {
	WhiteMoves []string
	BlackMoves []string
	WinnerRes *dvonn.MatchResult
}

func TestGame_02(t *testing.T) {
	input := "{\"WhiteMoves\":[\"d4\",\"g1\",\"f3\",\"h5\",\"b1\",\"e1\",\"f1\",\"g2\",\"g3\",\"d1\",\"h3\",\"f4\",\"i2\",\"e3\",\"h2\",\"b4\",\"a2\",\"j4\",\"f5\",\"e5\",\"j5\",\"i4\",\"d3\",\"i1\",\"j2\",\"f1:e1\",\"h5:h4\",\"j2:i2\",\"h4:f2\",\"a2:b3\",\"h2:g1\",\"i2:g2\",\"g3:g2\",\"g2:c2\",\"f2:i5\",\"i4:h3\",\"d1:c1\",\"f4:f5\",\"d3:e4\",\"j5:i5\",\"b4:b3\"]," +
		"\"BlackMoves\":[\"i5\",\"b2\",\"c2\",\"a3\",\"h1\",\"k5\",\"e2\",\"c5\",\"h4\",\"a1\",\"j3\",\"i3\",\"g4\",\"c1\",\"g5\",\"f2\",\"k4\",\"d5\",\"c4\",\"k3\",\"b3\",\"d2\",\"e4\",\"c3\",\"k5:j4\",\"k4:k3\",\"d5:e5\",\"h1:i1\",\"i1:k3\",\"c4:d4\",\"a1:b2\",\"d4:b2\",\"e5:e3\",\"i3:j4\",\"g4:f3\",\"j4:g1\",\"c3:c2\",\"d2:c1\",\"a3:b3\",\"g1:b1\"]," +
		"\"WinnerRes\":{\"WinningColor\":\"BLACK\",\"WinnerScore\":19,\"LoserScore\":9}}"
	gamePlayStr := &GamePlayStore{}
	json.Unmarshal([]byte(input), gamePlayStr)

	players := make([]dvonn.Player, 0)
	player1 := dvonn.GetPlayer("amit", "amit1234", dvonn.WHITE)
	player2 := dvonn.GetPlayer("kat", "1234amit", dvonn.BLACK)
	players = append(players, player1)
	players = append(players, player2)
	dvonnGame := dvonn.GetDvonnGame(players, player1)
	currentPlayer := player1

	whiteMoves := gamePlayStr.WhiteMoves
	blackMoves := gamePlayStr.BlackMoves
	expectedWinner := gamePlayStr.WinnerRes

	// start with white and go on playing the move for the next player
	whiteCounter, blackCounter := 0, 0

	for {
		moveIds := make([]string, 0)
		if dvonnGame.GetGamePhase() == dvonn.PLACEMENT_PHASE {
			if currentPlayer.GetPlayerColor() == dvonn.WHITE {
				moveIds = append(moveIds, strings.ToLower(whiteMoves[whiteCounter]))
				whiteCounter++
			} else {
				moveIds = append(moveIds, strings.ToLower(blackMoves[blackCounter]))
				blackCounter++
			}
		} else {
			if currentPlayer.GetPlayerColor() == dvonn.WHITE {
				moveIds = append(moveIds, strings.Split(strings.ToLower(whiteMoves[whiteCounter]), ":")...)
				whiteCounter++
			} else {
				moveIds = append(moveIds, strings.Split(strings.ToLower(blackMoves[blackCounter]), ":")...)
				blackCounter++
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
				whiteCounter -= 1
			} else {
				blackCounter -= 1
			}
		}
		currentPlayer = moveRes.GetNextPlayer()
	}

	winner, err := dvonnGame.GetGameWinner()
	if winner == nil || err != nil {
		// game should have ended
		t.Error("game should have ended by now")
	}
	if winner != nil && winner.GetWinnerColor() != expectedWinner.GetWinnerColor() {
		t.Error("black should have won")
	}

	if winner != nil && winner.LoserScore != expectedWinner.LoserScore {
		t.Error("loser score should be {}", expectedWinner.LoserScore)
	}

	if winner != nil && winner.WinnerScore != expectedWinner.WinnerScore {
		t.Error("winner score should be {}", expectedWinner.WinnerScore)
	}
}


