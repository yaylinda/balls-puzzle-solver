package models

import (
	"strings"
)

// getNextValidMoves returns the next valid moves for the board
func (b *Board) getNextValidMoves() []*Move {
	var moves []*Move

	for _, t1 := range b.Towers {
		for _, t2 := range b.Towers {
			t1 := t1
			t2 := t2
			move := &Move{From: t1, To: t2}
			if b.isMoveValid(move) {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

// isMoveValid returns if the given move is valid for the board
func (b *Board) isMoveValid(move *Move) bool {
	if move.From.Index == move.To.Index ||
		move.From.isEmpty() ||
		move.From.isComplete() ||
		move.To.isFull() {
		return false
	}

	ballToMove, _ := move.From.getTopBall()
	if ballToMove == nil {
		// This should not happen, since we already checked if the Tower is empty
		return false
	}

	ballToPutOnTopOf, _ := move.To.getTopBall()
	if ballToPutOnTopOf == nil {
		// The destination Tower is empty, so we can put any ball on top of it
		return true
	}

	// Verify if the ball to move has the same color as the ball to put on top of
	return ballToMove.Color == ballToPutOnTopOf.Color
}

// applyMove applies the move to the board and returns the new board
func (b *Board) applyMove(move *Move) *Board {
	newBoard := b.deepCopy()

	// Nil the ball in the "from" tower
	fromTowerIndex := move.From.Index
	fromBall, fromBallIndex := move.From.getTopBall()
	newBoard.Towers[fromTowerIndex].Balls[fromBallIndex] = nil

	// Set the ball in the "to" tower
	toTowerIndex := move.To.Index
	_, toBallIndex := move.To.getTopBall()
	newBoard.Towers[toTowerIndex].Balls[toBallIndex-1] = &Ball{
		ID:    fromBall.ID,
		Color: fromBall.Color,
	}

	return newBoard
}

// deepCopy returns a copy of the board as a new pointer
func (b *Board) deepCopy() *Board {
	var newTowers []*Tower

	for _, tower := range b.Towers {
		tower := tower
		var newBalls []*Ball
		for _, ball := range tower.Balls {
			ball := ball
			var newBall *Ball
			if ball == nil {
				newBall = nil
			} else {
				newBall = &Ball{ID: ball.ID, Color: ball.Color}
			}
			newBalls = append(newBalls, newBall)
		}
		newTowers = append(
			newTowers,
			&Tower{Balls: newBalls, Index: tower.Index},
		)
	}

	return &Board{Towers: newTowers}
}

// IsSolved returns if the board is solved
func (b *Board) IsSolved(expectedEmpty int) bool {
	var numEmptyTowers int
	var numCompleteTowers int

	for _, tower := range b.Towers {
		if tower.isEmpty() {
			numEmptyTowers++
			continue
		}
		if tower.isComplete() {
			numCompleteTowers++
			continue
		}
	}

	// fmt.Printf("\t\t numEmptyTowers=%d, numCompleteTowers=%d\n", numEmptyTowers, numCompleteTowers)

	return numEmptyTowers == expectedEmpty && numCompleteTowers == len(b.Towers)-expectedEmpty
}

// isEqual returns if the board's towers are equal to the other board's towers
func (b *Board) isEqual(other *Board) bool {
	for i, t := range b.Towers {
		if !t.isEqual(other.Towers[i]) {
			return false
		}
	}
	return true
}

// String method to return the string representation of a Board
func (b *Board) String() string {
	var builder strings.Builder

	for _, tower := range b.Towers {
		builder.WriteString("\t")
		builder.WriteString(tower.String())
		builder.WriteString("\n")
	}

	return builder.String()
}
