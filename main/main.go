package main

import (
	"../dvonn"
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

func main() {
	players := make([]dvonn.Player, 0)
	player1 := dvonn.GetPlayer("pushpender", "amit123", dvonn.WHITE)
	player2 := dvonn.GetPlayer("arman", "amit123", dvonn.BLACK)
	players = append(players, player1)
	players = append(players, player2)
	dvonnGame := dvonn.GetDvonnGame(players, player1)

	dvonnGame.Move(player1, "f1")
	dvonnGame.Move(player2, "g3")
}
