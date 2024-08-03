package models

import (
	"fmt"
	"strings"
)

// GetNextPossibleStates returns the next possible states from the current state
func (bs *BoardState) GetNextPossibleStates() []*BoardState {
	validMoves := bs.Board.getNextValidMoves()

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

// isEqual returns if the BoardState is equal to the other BoardState
func (bs *BoardState) isEqual(other *BoardState) bool {
	if !bs.Board.isEqual(other.Board) {
		return false
	}

	// Verify that the move is the same
	return bs.Move.To.Index == other.Move.To.Index && bs.Move.From.Index == other.Move.From.Index
}

// String returns the string representation of the BoardState
func (bs *BoardState) String() string {
	var builder strings.Builder

	builder.WriteString(bs.Board.String())
	if bs.Move != nil {
		builder.WriteString(fmt.Sprintf("%d->%d", bs.Move.From.Index, bs.Move.To.Index))
	}

	return builder.String()
}

// PrintSolution returns the string representation of the solution
func (bs *BoardState) PrintSolution() string {
	var builder strings.Builder

	builder.WriteString("============================================\n")
	builder.WriteString(fmt.Sprintf("Solution with %d moves\n", len(bs.Previous)))
	builder.WriteString("============================================\n")

	boards := append(bs.Previous, bs.Board)
	for i, board := range boards {
		builder.WriteString(fmt.Sprintf("Move %d:\n", i))
		builder.WriteString(board.String())
		builder.WriteString("----------------------------------\n")
	}

	return builder.String()
}
