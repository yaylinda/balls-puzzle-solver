package models

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
