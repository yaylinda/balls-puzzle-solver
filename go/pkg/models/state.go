package models

import "balls_puzzle_solver/pkg/utils"

// GetNextPossibleStates returns the next possible states from the current state
func (bs *BoardState) GetNextPossibleStates() []*BoardState {
	validMoves := bs.Board.getNextValidMoves()

	var nextStates []*BoardState

	for _, move := range validMoves {
		newBoard := bs.Board.applyMove(move)

		newBoardState := &BoardState{
			Board:    newBoard,
			Move:     move,
			Previous: bs,
		}

		nextStates = append(nextStates, newBoardState)
	}

	return nextStates
}

// IsSolved returns if the board state is solved
func (bs *BoardState) IsSolved(expectedEmpty int) bool {
	return bs.Board.isSolved(expectedEmpty)
}

// GetSolvedPath returns the path to the solved board
func (bs *BoardState) GetSolvedPath() []*Board {
	if !bs.IsSolved(2) {
		return nil
	}

	var path []*Board
	currentState := bs
	for currentState != nil {
		path = append(path, currentState.Board)
		currentState = currentState.Previous
	}

	utils.ReverseArray(path)
	return path
}

// Hash returns a unique "hashed" string representation of the board
func (bs *BoardState) Hash() string {
	return bs.Board.hash()
}
