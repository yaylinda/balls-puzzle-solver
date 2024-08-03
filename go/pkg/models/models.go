package models

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

// Board struct represents a collection of Towers
type Board struct {
	Towers []*Tower
}

// BoardState struct represents a board state node in the search tree
type BoardState struct {
	Board    *Board
	Move     *Move
	Previous []*Board
}
