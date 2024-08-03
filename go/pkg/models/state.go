package models

import (
	"fmt"
	"strings"
)

// GetNextPossibleStates returns the next possible states from the current state
func (bs *BoardState) GetNextPossibleStates() []*BoardState {
	validMoves := bs.Board.getNextValidMoves()

	// TODO: Trimming logic

	var nextStates []*BoardState

	for _, move := range validMoves {
		newBoard := bs.Board.applyMove(move)

		newBoardState := &BoardState{
			Board:    newBoard,
			Move:     move,
			Previous: append(bs.Previous, bs.Board),
		}

		nextStates = append(nextStates, newBoardState)
	}

	return nextStates
}

// IsSolved returns if the board state is solved
func (bs *BoardState) IsSolved(expectedEmpty int) bool {
	return bs.Board.isSolved(expectedEmpty)
}

// Hash returns a unique "hashed" string representation of the board
func (bs *BoardState) Hash() string {
	return bs.Board.hash()
}

// PrintSolution returns the string representation of the solution
func (bs *BoardState) PrintSolution() string {
	var builder strings.Builder

	builder.WriteString("============================================\n")
	builder.WriteString(
		fmt.Sprintf(
			"Solution with %d moves\n",
			len(bs.Previous),
		),
	)
	builder.WriteString("============================================\n")

	boards := append(bs.Previous, bs.Board)
	for i, board := range boards {
		builder.WriteString(fmt.Sprintf("Move %d:\n", i))
		builder.WriteString(board.string())
		builder.WriteString("----------------------------------\n")
	}

	return builder.String()
}
