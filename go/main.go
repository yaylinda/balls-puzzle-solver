package main

func main() {
	// Example usage

	// Creating a new ball
	redBall := &Ball{Color: "red"}

	// Creating a new move
	move := &Move{
		From: 0,
		To:   1,
		Ball: redBall,
	}

	// Creating a new board state
	board := make([][]*Ball, 3)
	for i := range board {
		board[i] = make([]*Ball, 3)
	}

	boardState := &BoardState{
		Board: board,
		Move:  move,
	}

	// Example to set a ball in the board
	boardState.Board[0][0] = redBall

	// Print the color of the ball at board[0][0]
	if boardState.Board[0][0] != nil {
		println("Ball color at board[0][0]:", boardState.Board[0][0].Color)
	}

	// Print the move details
	if boardState.Move != nil {
		println(
			"Move from:",
			boardState.Move.From,
			"to:",
			boardState.Move.To,
			"with ball color:",
			boardState.Move.Ball.Color,
		)
	}
}
