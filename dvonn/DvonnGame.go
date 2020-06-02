package dvonn

import (
	"errors"
	"log"
)

type DvonnGame struct {
	board *DvonnBoard
	players []Player
	currentTurn Player
	gamePhase GamePhase
	unusedChips int
}

func GetDvonnGame(players []Player, whitePlayer Player) *DvonnGame {
	return &DvonnGame {
		board:GetDvonnBoard(),
		players:players,
		currentTurn:whitePlayer,
		gamePhase:PLACEMENT_PHASE,
		unusedChips:49,
	}
}

func (dg *DvonnGame) togglePlayer() {
	if dg.currentTurn.GetPlayerColor() == WHITE {
		dg.currentTurn = dg.GetPlayerByColor(BLACK)
	} else {
		dg.currentTurn = dg.GetPlayerByColor(WHITE)
	}
}

func (dg *DvonnGame) GetPlayerByColor(color ChipColor) Player {
	player := Player{}  // dummy initialization
	for _, player := range dg.GetPlayers() {
		if player.GetPlayerColor() == color {
			return player;
		}
	}
	log.Fatal("[DvonnGame.GetPlayerByColor] can not find player with color " + color)
	return player
}


func (dg *DvonnGame) GetPlayers() []Player {
	return dg.players
}

/*
 NOTE: Ideally the players needs to be passed at the time when Instance of Game is being created, adding player
	   in b/w doesn't make any sense unless the scope of game changes in future.
 */
func (dg *DvonnGame) AddPlayer(player Player) {
	dg.players = append(dg.players, player)
}

// TODO: this method will not be used extensively, very rarely this has some use case but don't forget to write testcase
func (dg *DvonnGame) RemovePlayer(playerId string) {
	remIndex := -1
	for index, player := range dg.players {
		if playerId == player.GetId() {
			remIndex = index
		}
	}
	if remIndex == -1 {
		return
	}
	dg.players = append(dg.players[:remIndex], dg.players[remIndex+1:]...)
}


/*
 NOTE: this shall be majorly be used for validating the user move, i.e. only current player turn can take move
 Also make sure that the current player is getting updated after each "valid" move.
 */
func (dg *DvonnGame) GetCurrentPlayer() Player {
	return dg.currentTurn
}


func (dg *DvonnGame) GetGamePhase() GamePhase {
	if dg.unusedChips > 0 {
		return PLACEMENT_PHASE
	}
	return MOVEMENT_PHASE
}


/*
 There will be two types of movements:
	a) Placement type: i.e. only player and position will be required.
	b) Movement type: along with the player the origin position and destination position is required

 The client will call this method with variadic parameters, if the Game phase is PLACEMENT phase we expect the client
 to pass player name and placeDestination Id
 Otherwise, we expect `player`, `originId`, `destinationId` as the arguments, on our side we will anyway validate
 the passed arguments and return error code if invalid arguments are passed.

 NOTE: the client should not pass the chips, depending on the player we place the chips in placement phase and in
 movement phase we check the origin position and destination position and get chips from accordingly.
 */
func (dg *DvonnGame) Move(player Player, paths ...string) (bool, error) {
	if dg.GetGamePhase() == PLACEMENT_PHASE {
		if len(paths) > 1 {
			return false, errors.New("can not get multiple paths as the game is still in PLACEMENT phase")
		}
		destId := paths[0]
		placeValid, err := dg._canPlace(player, destId)
		if !placeValid || err != nil {
			return false, err
		}
		dg.board.Fill(destId, []Chip{{player.GetPlayerColor()}})
		dg.unusedChips = dg.unusedChips - 1;  // updating count of unused chips as chips are now being involved in game
		return true, nil
	} else if dg.GetGamePhase() == MOVEMENT_PHASE {
		if len(paths) != 2 {
			return false, errors.New("only two ids are required in movement phase, which are origin and dest. id")
		}
		orgId := paths[0]
		dstId := paths[1]
		moveValid, err := dg._canMove(player, orgId, dstId)
		if !moveValid || err != nil {
			return false, err
		}
		// do movement
		origStack := dg.board.GetCells()[orgId].GetChipsStack();
		dg.board.DeFill(orgId)
		dg.board.Fill(dstId, origStack)
		dg.board.RemoveDisconnectedCells()
	} else {
		log.Fatal("GamePhase not available")
	}
	dg.togglePlayer()
	return false, nil  // anyway this is unreachable code
}

func (dg *DvonnGame) _canPlace(player Player, destId string) (bool, error) {
	isValid := true
	errM := ""
	// validate player turn
	isValid = dg.isPlayerTurnValid(player)
	if !isValid {
		errM = "another player is required to play the move"
	}

	// validate destination place
	isValid = dg.board.IsPlaceVacant(destId)
	if !isValid {
		errM = "position " + destId + "is already occupied"
	}

	if errM == "" {
		return isValid, nil
	}
	return isValid, errors.New(errM)
}

func (dg *DvonnGame) _canMove(player Player, originId, destId string) (bool, error) {
	isValid := true
	errM := ""
	// validate game phase
	if isValid = dg.GetGamePhase() == MOVEMENT_PHASE; !isValid {
		errM = "game phase is not valid, it should be movement phase"
	}
	// validate player turn
	if isValid = dg.isPlayerTurnValid(player); !isValid {
		errM = "different player is required to play the move"
	}
	// validate: origin and destination place should not be empty
	if isValid = dg.board.GetCells()[originId].IsEmpty() && dg.board.GetCells()[destId].IsEmpty(); !isValid {
		errM = "the origin aid destination place can not be empty"
	}
	if isValid = dg.board.GetCells()[originId].IsEmpty(); !isValid {
		errM = "the origin place can not be empty"
	}
	if isValid = dg.board.GetCells()[destId].IsEmpty(); !isValid {
		errM = "the destination place can not be empty"
	}
	// validate source destination by calculating adjacent nodes i.e. distance should be length of stack on origin node
	isDestValid := false
	possibleDestNodes := dg.board.GetCells()[originId].GetStraightAdjacentOnLevel(dg.board.GetCells()[originId].GetStackLength())
	for _, node := range possibleDestNodes {
		if destId == node.GetIdentifier() {
			isDestValid = true
		}
	}
	if isValid = isDestValid; !isValid {
		errM = "not possible to place stack from " + originId + " to " + destId
	}
	return isValid, errors.New(errM)
}

// validate player, i.e. only the current player should be able to make movement
func (dg *DvonnGame) isPlayerTurnValid(player Player) bool {
	if player.GetId() != dg.currentTurn.GetId() {
		return false
	}
	return true
}

