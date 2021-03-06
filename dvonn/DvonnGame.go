package dvonn

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"log"
)

type DvonnGame struct {
	board       *DvonnBoard
	players     []Player
	currentTurn Player
	gamePhase   GamePhase
	unusedChips int

	// some props for undo/redo operation
	HEAD int  // an index pointer which shall be pointing to the current game state

	// to storing game state, we shalling be storing the moves played and then simulate those steps
	movesPlayed [][]string
}

func GetDvonnGame(players []Player, whitePlayer Player) *DvonnGame {
	return &DvonnGame{
		board:          GetDvonnBoard(),
		players:        players,
		currentTurn:    whitePlayer,
		gamePhase:      PLACEMENT_PHASE,
		unusedChips:    49,
		movesPlayed:    make([][]string, 0),
		HEAD:           -1,
	}
}

func (dg *DvonnGame) GetBoard() *DvonnBoard {
	return dg.board
}

func (dg *DvonnGame) togglePlayer() {
	dg.currentTurn = dg.getOppositeColoredPlayer(dg.currentTurn)
}

func (dg *DvonnGame) getOppositeColoredPlayer(player Player) Player {
	if player.GetPlayerColor() == WHITE {
		return dg.GetPlayerByColor(BLACK)
	} else {
		return dg.GetPlayerByColor(WHITE)
	}
}

func (dg *DvonnGame) GetPlayerByColor(color ChipColor) Player {
	player := Player{} // dummy initialization
	for _, player := range dg.GetPlayers() {
		if player.GetPlayerColor() == color {
			return player
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
 This is a wrapper method on _move method which is actually doing all the logic related to move.
 NOTE: this method is purposely written to be consumed by game client for playing moves
 This method will also serve the purpose of writing logic on result obtained from the last move played.
 This method should be called once for one move.
*/
func (dg *DvonnGame) Move(player Player, paths ...string) MoveResult {
	moveRes := dg.move(player, paths...)
	// based on results do operation like set next player as current player, etc
	if moveRes.isGameOver {
		return moveRes
	}
	if moveRes.isActionSuccess {
		nextPlayer, _ := dg.GetNextTurnPlayer()
		moveRes.SetNextPlayer(nextPlayer)
		// balance out the HEAD pointer and game state stack, why is this balancing required?
		// coz, if last operation was undo or redo, then HEAD pointer must be not in sync with the movesPlayed, so
		// we need to remove the moves played from this stack and undo operation has been done
		if len(dg.movesPlayed) > 0 {
			dg.movesPlayed = dg.movesPlayed[:dg.HEAD+1]
		}
		dg.HEAD++
		moves := make([]string, 0)
		moves = append(moves, paths...)
		dg.movesPlayed = append(dg.movesPlayed, moves)
	} else {
		// validation failure occurred, so allow the same player to make the same move as the game isn't finished yet
		moveRes.SetNextPlayer(player)
	}
	// if the placement phase has just ended then the next move should be done by WHITE
	if moveRes.gamePhase == PLACEMENT_PHASE && dg.unusedChips == 0 { // NOTE: fetch phase from moreRes only for comparing, this should be the source of truth
		moveRes.SetNextPlayer(dg.GetPlayerByColor(WHITE))
	}
	// set the next player as the current player so that the validation could be done when the client is calling the
	// move method again
	dg.currentTurn = moveRes.GetNextPlayer()
	return moveRes
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
// TODO: send in response a set of actions that the implementer could perform
// like MOVEMENT_DONE, PLACEMENT_DONE, NO_VALID_MOVE(return along with the first two stages), GAME_END_STATE,
// WINNER, etc, think more on it
func (dg *DvonnGame) move(player Player, paths ...string) MoveResult {
	if dg.GetGamePhase() == PLACEMENT_PHASE {
		if len(paths) > 1 {
			errM := "can not get multiple paths as the game is still in PLACEMENT phase"
			return GetMoveResult(false, false,
				ERROR_ARGUMENT_COUNT_MISMATCH, errM, errors.New(errM), PLACEMENT_PHASE)
		}
		destId := paths[0]
		validationRes := dg._canPlace(player, destId)
		if validationRes.err != nil {
			return validationRes
		}
		// the first 3 moves should have to be red colors only
		if dg.unusedChips > 46 {
			dg.board.Fill(destId, []Chip{{RED}})
		} else {
			dg.board.Fill(destId, []Chip{{player.GetPlayerColor()}})
		}
		dg.unusedChips = dg.unusedChips - 1 // updating count of unused chips as chips are now being involved in game
		return GetMoveResult(false, true, ERROR_UNKNOWN, "", nil, PLACEMENT_PHASE)
	} else if dg.GetGamePhase() == MOVEMENT_PHASE {
		if len(paths) != 2 {
			errM := "only two ids are required in movement phase, which are origin and destination id"
			return GetMoveResult(dg.IsGameOver(), false,
				ERROR_ARGUMENT_COUNT_MISMATCH, errM, errors.New(errM), MOVEMENT_PHASE)
		}
		orgId := paths[0]
		dstId := paths[1]
		validationRes := dg._canMove(player, orgId, dstId)
		if validationRes.err != nil {
			return validationRes
		}
		// do movement
		origStack := dg.board.GetCells()[orgId].GetChipsStack()
		dg.board.DeFill(orgId)
		dg.board.Fill(dstId, origStack)
		dg.board.RemoveDisconnectedCells()
	} else {
		log.Fatal("GamePhase not available")
	}
	return GetMoveResult(dg.IsGameOver(), true,
		ERROR_UNKNOWN, "", nil, MOVEMENT_PHASE)
}

func (dg *DvonnGame) _canPlace(player Player, destId string) MoveResult {
	isValid := true
	errM := ""
	// validate player turn
	isValid = dg.isPlayerTurnValid(player)
	if !isValid {
		errM = "another player is required to play the move"
		return GetMoveResult(false, false,
			ERROR_INVALID_PLAYER_TURN, errM, errors.New(errM), PLACEMENT_PHASE)
	}

	// validate destination place
	isValid = dg.board.IsPlaceVacant(destId)
	if !isValid {
		errM = "position " + destId + "is already occupied"
		return GetMoveResult(false, false, ERROR_DESTINATION_ALREADY_OCCUPIED, errM,
			errors.New(errM), PLACEMENT_PHASE)
	}

	// all validations are passed
	return GetMoveResult(false, true, ERROR_UNKNOWN,
		errM, nil, PLACEMENT_PHASE)
}

/*
 player in this method is the current player who is making the move in the game.
 NOTE: the order of validation is important below, eg: validating if node at origin id should
 		have same color as that of the current player playing move should be done after we are
		done validating that originId/destinationId is valid (i.e. non-empty and have free adjacent)
 Violating the order wouldn't always impact, but for proper validation message to client, it gets necessary
 to maintain order.
 */
func (dg *DvonnGame) _canMove(player Player, originId, destId string) MoveResult {
	isValid := true
	errM := ""
	// validate game phase
	if isValid = dg.GetGamePhase() == MOVEMENT_PHASE; !isValid {
		errM = "game phase is not valid, it should be movement phase"
		return GetMoveResult(dg.IsGameOver(), false,
			ERROR_INVALID_GAME_PHASE, errM, errors.New(errM), MOVEMENT_PHASE)
	}
	// validate player turn
	if isValid = dg.isPlayerTurnValid(player); !isValid {
		errM = "different player is required to play the move"
		return GetMoveResult(dg.IsGameOver(), false,
			ERROR_INVALID_PLAYER_TURN, errM, errors.New(errM), MOVEMENT_PHASE)
	}
	// validate: origin and destination place should not be empty
	if isValid = !dg.board.IsCellEmpty(originId) && !dg.board.IsCellEmpty(destId); !isValid {
		errM = "the origin and destination place can not be empty"
		return GetMoveResult(dg.IsGameOver(), false,
			ERROR_EMPTY_ORIGIN_DESTINATION, errM, errors.New(errM), MOVEMENT_PHASE)
	}
	if isValid = !dg.board.IsCellEmpty(originId); !isValid {
		errM = "the origin place can not be empty"
		return GetMoveResult(dg.IsGameOver(), false,
			ERROR_EMPTY_ORIGIN, errM, errors.New(errM), MOVEMENT_PHASE)
	}
	if isValid = !dg.board.IsCellEmpty(destId); !isValid {
		errM = "the destination place can not be empty"
		return GetMoveResult(dg.IsGameOver(), false,
			ERROR_EMPTY_DESTINATION, errM, errors.New(errM), MOVEMENT_PHASE)
	}

	// validate: only those chips can be moved where at-least on of the adjacent is free.
	if isValid = dg.board.GetCells()[originId].HasFreeEdge(); !isValid {
		errM = "only chips with some free surroundings can be moved origin id: " + originId
		return GetMoveResult(dg.IsGameOver(), false,
			ERROR_NO_FREE_ADJACENT_FOUND, errM, errors.New(errM), MOVEMENT_PHASE)
	}

	// validate source destination by calculating adjacent nodes i.e. distance should be length of stack on origin node
	isDestValid := false
	possibleDestNodes := dg.board.GetPossibleMoveFor(originId)
	for _, node := range possibleDestNodes {
		if destId == node.GetIdentifier() {
			isDestValid = true
		}
	}
	if isValid = isDestValid; !isValid {
		errM = "not possible to place stack from " + originId + " to " + destId
		return GetMoveResult(dg.IsGameOver(), false,
			ERROR_INVALID_DESTINATION_SELECTED, errM, errors.New(errM), MOVEMENT_PHASE)
	}

	// validate if origin node has same color as that of the current turn player color
	topChip, _ := dg.GetBoard().GetCells()[originId].GetTopChips()
	if isValid = topChip.GetColor() == player.GetPlayerColor(); !isValid {
		errM = "Player with "+string(player.GetPlayerColor())+" can not move chip with "+string(topChip.GetColor())+" color"
		return GetMoveResult(dg.IsGameOver(), false,
			ERROR_WRONG_ORIGIN_SELECTED, errM, errors.New(errM), MOVEMENT_PHASE)
	}


	// all validation are passed
	return GetMoveResult(dg.IsGameOver(), true,
		ERROR_UNKNOWN, "", nil, MOVEMENT_PHASE)
}

// validate player, i.e. only the current player should be able to make movement
func (dg *DvonnGame) isPlayerTurnValid(player Player) bool {
	return player.GetPlayerColor() == dg.currentTurn.GetPlayerColor()
}

/*
 Checks if any valid move is left for a player
 arg: player for whom checking valid move is required
 return: return true if any valid move is available, else false

 Solution:
	- get all cell ids which have the respective player color on top of stack
	- If any possible move from that cell id is present then return true
	- If no possible move could be found then return false, stating no valid move left
*/
func (dg *DvonnGame) IsValidMoveLeftForPlayer(playerColor ChipColor) bool {
	cellIds := dg.board.GetCellIdsByStackColor(playerColor)
	// for each id check if possible movement can be done from that stack
	for _, id := range cellIds {
		if len(dg.board.GetPossibleMoveFor(id)) > 0 && dg.board.GetCells()[id].HasFreeEdge() {
			return true
		}
	}
	return false
}

/*
 tells if the other player has some valid move left to play or not
*/
func (dg *DvonnGame) HasNextPlayerAValidMove(currentPlayer Player) bool {
	// in placement phase all players will always have a valid move to play
	if dg.GetGamePhase() == PLACEMENT_PHASE {
		return true
	}
	currentPlayColor := currentPlayer.GetPlayerColor()
	var nextPlayerColor ChipColor
	if currentPlayColor == WHITE {
		nextPlayerColor = BLACK
	} else {
		nextPlayerColor = WHITE
	}
	return dg.IsValidMoveLeftForPlayer(nextPlayerColor)
}

/*
 returns false if any one of the player has any move left to play in MOVEMENT phase, else return true
*/
func (dg *DvonnGame) IsGameOver() bool {
	// NOTE that just when all the chips are placed from that point the game will transition into movement phase
	return dg.GetGamePhase() != PLACEMENT_PHASE && !dg.IsValidMoveLeftForPlayer(BLACK) && !dg.IsValidMoveLeftForPlayer(WHITE)
}

/*
 Returns winner color
*/
func (dg *DvonnGame) GetGameWinner() (*MatchResult, error) {
	var winnerColor WinnerColor
	winnerScore := 0
	loserScore := 0
	if dg.IsGameOver() {
		playerOnePieces := dg.board.GetCountOfPiecesControlledBy(dg.GetPlayers()[0].GetPlayerColor())
		playerTwoPieces := dg.board.GetCountOfPiecesControlledBy(dg.GetPlayers()[1].GetPlayerColor())
		if playerOnePieces > playerTwoPieces {
			winnerColor = GetWinnerColorFromPlayerColor(dg.GetPlayers()[0].GetPlayerColor())
			winnerScore = playerOnePieces
			loserScore = playerTwoPieces
		} else if playerOnePieces < playerTwoPieces {
			winnerColor = GetWinnerColorFromPlayerColor(dg.GetPlayers()[1].GetPlayerColor())
			winnerScore = playerTwoPieces
			loserScore = playerOnePieces
		} else {
			winnerColor = WINNER_DRAW
			winnerScore, loserScore = playerOnePieces, playerTwoPieces
		}
		return &MatchResult{LoserScore: loserScore, WinnerScore:winnerScore, WinningColor:winnerColor}, nil
	}
	return &MatchResult{}, errors.New("this game is not over yet")

}

/*
 Returns the player whose turn is next, i.e. after the current move
 Logic: if game is not over then either of the player has a valid move left
	so, it returns other player if other player has a valid move, otherwise
	it returns the current player itself
*/
func (dg *DvonnGame) GetNextTurnPlayer() (Player, error) {
	var player Player
	if dg.IsGameOver() {
		return player, errors.New("game is already over")
	}
	if dg.HasNextPlayerAValidMove(dg.GetCurrentPlayer()) {
		return dg.getOppositeColoredPlayer(dg.GetCurrentPlayer()), nil
	}
	// since no valid move is left for the other player, so current player should continue playing
	return dg.GetCurrentPlayer(), nil
}


/*
 Usage of this method: this method doesn't mutate the current state rather, it will return one state which will be
 the previous state. So, the caller needs to make sure that instance is being reassigned to what is being returned from hear

 In case of consecutive undo operation:
	method checks if HEAD pointer is pointing to 0, then it can't perform undo operation and this method will return
	a flag to indicate that

 Each time we perform a successful undo operation we decrement the HEAD pointer
 And on any successful move post undo operation we pop off the top elements after HEAD pointer from the movesPlayed stack
 to balance the game state stack and HEAD pointer.
 */
func (dg *DvonnGame) Undo() (*DvonnGame, bool) {
	if dg.HEAD <= 0 || len(dg.movesPlayed) <= 0 || dg.HEAD > len(dg.movesPlayed) {
		return nil, false
	}
	// if game is over, undo should not be applicable
	if dg.IsGameOver() {
		return nil, false
	}
	dg.HEAD--
	// till HEAD pointer simulate the game and return that game instance
	stepsToSimulate := 0
	newGameInstance := GetDvonnGame(dg.GetPlayers(), dg.GetPlayerByColor(WHITE))
	for stepsToSimulate <= dg.HEAD {
		newGameInstance.Move(newGameInstance.GetCurrentPlayer(), dg.movesPlayed[stepsToSimulate]...)
		stepsToSimulate++
	}
	return newGameInstance, true
}


/*
 Usage of this method: It doesn't mutate the game state, for that the caller need to make sure for reinitializing
 the game state with what is returned from this method.

 NOTE: redo can only be performed when last step would be undo operation.
 So, there is a clear validation for the HEAD pointer must be smaller than the movesPlayed stack length.

 if undo would have been done in last step then dg.HEAD would be lesser than the length of movesPlayed stack
 coz, an decrement for the HEAD must have happened
 if undo wasn't done in the last step then the length of the game state stack would match the HEAD (NOTE:
 in case of normal play steps I do pop off the extra top elements above the HEAD to maintain consistency)
 */
func (dg *DvonnGame) Redo() bool {

	if dg.HEAD >= len(dg.movesPlayed) || len(dg.movesPlayed) == 0 || dg.HEAD == 0 {
		return false
	}
	dg.HEAD++
	dg.Move(dg.GetCurrentPlayer(), dg.movesPlayed[dg.HEAD]...)
	return true
}

// Clone deep-copies a to b
func Clone(a, b interface{}) {

	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(a)
	dec.Decode(b)
}

// DeepCopy deepcopies a to b using json marshaling
func DeepCopy(a, b interface{}) {
	byt, _ := json.Marshal(a)
	json.Unmarshal(byt, b)
}
