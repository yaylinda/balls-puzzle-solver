package models

import "strings"

// getTopBall returns the top ball of the Tower, and the index of the ball.
// If the Tower is empty, it returns nil and the length of the Balls slice
func (t *Tower) getTopBall() (*Ball, int) {
	for i := range t.Balls {
		if t.Balls[i] != nil {
			return t.Balls[i], i
		}
	}
	return nil, len(t.Balls)
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

// isComplete returns if the Tower is full and all balls are the same color
func (t *Tower) isComplete() bool {
	if !t.isFull() {
		return false
	}

	color := t.Balls[0].Color

	for i := range t.Balls {
		if t.Balls[i].Color != color {
			return false
		}
	}

	return true
}

// isEqual returns if the tower's balls are equal to the other tower's balls at each position
func (t *Tower) isEqual(other *Tower) bool {
	for i := range t.Balls {
		if t.Balls[i] == nil && other.Balls[i] == nil {
			continue
		}
		if t.Balls[i] == nil || other.Balls[i] == nil {
			return false
		}
		if t.Balls[i].ID != other.Balls[i].ID {
			return false
		}
	}

	return true
}

// String method to return the string representation of a Tower
func (t *Tower) String() string {
	var builder strings.Builder

	builder.WriteString("[")
	for i, ball := range t.Balls {
		if i > 0 {
			builder.WriteString(",")
		}
		if ball == nil {
			builder.WriteString(" ")
		} else {
			builder.WriteString(ball.Color)
		}
	}
	builder.WriteString("]")

	return builder.String()
}
