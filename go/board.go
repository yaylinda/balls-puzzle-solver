package main

import (
	"fmt"
	"strings"
)

// Ball struct represents a ball with a color
type Ball struct {
	ID    string
	Color string
}

// Move struct represents moving the top ball from one Tower to another Tower
type Move struct {
	From *Tower
	To   *Tower
}

// Tower struct represents a collection of 4 balls on top of each other
type Tower struct {
	Index int
	Balls []*Ball
}

// getTopBall returns the top ball of the Tower, and the index of the ball.
// If the Tower is empty, it returns nil and -1
func (t *Tower) getTopBall() (*Ball, int) {
	for i := range t.Balls {
		if t.Balls[i] != nil {
			return t.Balls[i], i
		}
	}
	return nil, -1
}

// isEmpty returns if the Tower is empty
func (t *Tower) isEmpty() bool {
	for i := range t.Balls {
		if t.Balls[i] != nil {
			return false
		}
	}
	return true
}

// isFull returns if the Tower is full
func (t *Tower) isFull() bool {
	for i := range t.Balls {
		if t.Balls[i] == nil {
			return false
		}
	}
	return true
}

// BoardState struct represents the state of the board
type BoardState struct {
	Board []*Tower
	Move  Move
}

// applyMove returns a pointer to a new BoardState after applying the move
func (bs *BoardState) applyMove() *BoardState {
	if !bs.isMoveValid() {
		return nil
	}

	// Deep copy the towers and balls arrays
	var newTowers []*Tower
	for _, tower := range bs.Board {
		var newBalls []*Ball
		for _, ball := range tower.Balls {
			if ball == nil {
				newBalls = append(newBalls, nil)
				continue
			}
			newBalls = append(newBalls, &Ball{ID: ball.ID, Color: ball.Color})
		}
		newTowers = append(newTowers, &Tower{Balls: newBalls, Index: tower.Index})
	}

	// Create a new BoardState with the new towers copy
	newBoardState := &BoardState{
		Board: newTowers,
		Move:  bs.Move,
	}

	// Nil the ball in the "from" tower
	fromTowerIndex := bs.Move.From.Index
	fromBall, fromBallIndex := bs.Move.From.getTopBall()
	newBoardState.Board[fromTowerIndex].Balls[fromBallIndex] = nil

	// Set the ball in the "to" tower
	toTowerIndex := bs.Move.To.Index
	_, toBallIndex := bs.Move.To.getTopBall()
	newBoardState.Board[toTowerIndex].Balls[toBallIndex] = &Ball{
		ID:    fromBall.ID,
		Color: fromBall.Color,
	}

	return newBoardState
}

// isMoveValid returns if the given move is valid for the board
func (bs *BoardState) isMoveValid() bool {
	if bs.Move.From.isEmpty() {
		return false
	}

	if bs.Move.To.isFull() {
		return false
	}

	ballToMove, _ := bs.Move.From.getTopBall()
	if ballToMove == nil {
		// This should not happen, since we already checked if the Tower is empty
		return false
	}

	ballToPutOnTopOf, _ := bs.Move.To.getTopBall()
	if ballToPutOnTopOf == nil {
		// The destination Tower is empty, so we can put any ball on top of it
		return true
	}

	// Verify if the ball to move has the same color as the ball to put on top of
	return ballToMove.Color == ballToPutOnTopOf.Color
}

// isEqual compares the current BoardState with another BoardState
func (bs *BoardState) isEqual(other *BoardState) bool {
	// Check that the balls are in the same position
	for i := range bs.Board {
		for j := range bs.Board[i].Balls {
			if bs.Board[i].Balls[j] == nil && other.Board[i].Balls[j] == nil {
				continue
			}
			if bs.Board[i].Balls[j] == nil || other.Board[i].Balls[j] == nil {
				return false
			}
			if bs.Board[i].Balls[j].ID != other.Board[i].Balls[j].ID {
				return false
			}
		}
	}

	// Check that the move is the same
	return bs.Move.To.Index == other.Move.To.Index && bs.Move.From.Index == other.Move.From.Index
}

// String method to return the string representation of a BoardState
func (bs *BoardState) String() string {
	var builder strings.Builder

	for _, tower := range bs.Board {
		builder.WriteString("[")
		for i, ball := range tower.Balls {
			if i > 0 {
				builder.WriteString(",")
			}
			builder.WriteString(ball.ID)
		}
		builder.WriteString("]")
	}

	builder.WriteString(fmt.Sprintf("::%d->%d", bs.Move.From.Index, bs.Move.To.Index))

	return builder.String()
}
