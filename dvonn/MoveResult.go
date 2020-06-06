package dvonn

/*
 A type for returning the details of action taken on a player Move,
 The move could be played in both the phases PLACEMENT and MOVEMENT phase.
*/
type MoveResult struct {
	isGameOver      bool
	isActionSuccess bool
	errorCode       ErrorCode
	errorMessage    string
	err             error
	gamePhase       GamePhase
	nextPlayer      Player
}

func GetMoveResult(isGameOver, isActionSuccess bool,
	code ErrorCode, errM string, err error, phase GamePhase) MoveResult {
	return MoveResult{
		isGameOver:      isGameOver,
		isActionSuccess: isActionSuccess,
		errorCode:       code,
		errorMessage:    errM,
		err:             err,
		gamePhase:       phase,
	}
}

func (mr *MoveResult) SetNextPlayer(player Player) {
	mr.nextPlayer = player
}

// represents error states
type ErrorCode int

const (
	ERROR_UNKNOWN ErrorCode = iota
	ERROR_ARGUMENT_COUNT_MISMATCH
	ERROR_INVALID_PLAYER_TURN
	ERROR_DESTINATION_ALREADY_OCCUPIED

	ERROR_INVALID_GAME_PHASE
	ERROR_EMPTY_ORIGIN_DESTINATION
	ERROR_EMPTY_ORIGIN
	ERROR_EMPTY_DESTINATION
	ERROR_NO_FREE_ADJACENT_FOUND
	ERROR_INVALID_DESTINATION_SELECTED
)

func (mr *MoveResult) IsGameOver() bool {
	return mr.isGameOver
}

func (mr *MoveResult) GetNextPlayer() Player {
	return mr.nextPlayer
}

func (mr *MoveResult) IsActionSuccess() bool {
	return mr.isActionSuccess
}

func (mr *MoveResult) GetErrorCode() ErrorCode {
	return mr.errorCode
}

func (mr *MoveResult) GetErrorMessage() string {
	return mr.errorMessage
}
