package testcases

import (
	"../dvonn"
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
	dvonnGame.Move(player1, "f1")
	dvonnGame.Move(player2, "g3")

}
