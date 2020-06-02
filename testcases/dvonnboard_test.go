package testcases

import (
	"../dvonn"
	"testing"
)


func TestHexBoardTraversal_01(t *testing.T) {
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
	dvonnBoard := dvonn.GetDvonnBoard()
	dvonnBoard.Fill("a1", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("b1", []dvonn.Chip{dvonn.GetChips(dvonn.RED)})
	dvonnBoard.Fill("c1", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("d1", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("e1", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("f1", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("g1", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("h1", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("i1", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})

	dvonnBoard.Fill("a2", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("b2", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("c2", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("d2", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("e2", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("f2", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("g2", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("h2", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("i2", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("j2", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})

	dvonnBoard.Fill("a3", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("b3", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("c3", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("d3", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("e3", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("f3", []dvonn.Chip{dvonn.GetChips(dvonn.RED)})
	dvonnBoard.Fill("g3", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("h3", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("i3", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("j3", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("k3", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})

	dvonnBoard.Fill("b4", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("c4", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("d4", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("e4", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("f4", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("g4", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("h4", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("i4", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("j4", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("k4", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})

	dvonnBoard.Fill("c5", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("d5", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("e5", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("f5", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("g5", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("h5", []dvonn.Chip{dvonn.GetChips(dvonn.WHITE)})
	dvonnBoard.Fill("i5", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})
	dvonnBoard.Fill("j5", []dvonn.Chip{dvonn.GetChips(dvonn.RED)})
	dvonnBoard.Fill("k5", []dvonn.Chip{dvonn.GetChips(dvonn.BLACK)})

	disconnectedCells := dvonnBoard.GetDisconnectedCells()
	// disconnectedCells should be null
	if disconnectedCells != nil {
		t.Error("disconnected cells should be none")
	}

	/*
	         5
	      4   \ __
	   3   \ __/w \__
	    \ __/b \__/b \__
	  2  /b \__/b \__/b \__
	   \ \__/b \__/w \__/  \__
	  1  /w \__/b \__/  \__/w \__
	   \ \__/  \__/  \__/w \__/w \__
	     /  \__/  \__/w \__/b \__/b \__
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
	dvonnBoard.DeFill("a1")
	dvonnBoard.DeFill("b2")
	dvonnBoard.DeFill("c2")
	dvonnBoard.DeFill("d3")
	dvonnBoard.DeFill("e4")
	dvonnBoard.DeFill("f5")

	disconnectedCells = dvonnBoard.GetDisconnectedCells()
	if len(disconnectedCells) != 16 {
		t.Error("disconnected cells count should be 16 as stated in the above diagram")
	}
}