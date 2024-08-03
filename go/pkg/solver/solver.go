package solver

import (
	"fmt"

	"balls_puzzle_solver/pkg/models"
	"balls_puzzle_solver/pkg/utils"
)

type PrintOpts struct {
	FindShortest bool
}

// Solve solves the puzzle and returns the final state with all previous board
// states, and the number of iterations
func Solve(puzzle [][]string, opts PrintOpts) (*models.BoardState, int) {
	// Create a board from the puzzle
	initialBoard := utils.CreateBoardFromPuzzle(puzzle)

	// Keep track of the visited states
	visited := make(map[string]bool)
	willVisit := make(map[string]bool)

	// Create a queue with the initial state
	queue := []*models.BoardState{
		{
			Board:    initialBoard,
			Move:     nil,
			Previous: []*models.Board{},
		},
	}

	iteration := 1
	var shortestState *models.BoardState

	// Iterate over the queue
	for len(queue) > 0 {
		// Get the first element of the queue
		currentState := queue[0]

		// Check if the current state is the final state
		if currentState.IsSolved(2) {
			if !opts.FindShortest {
				return currentState, iteration
			}
			fmt.Printf(
				"\t[iter=%d] found solution with %d moves\n",
				iteration,
				len(currentState.Previous),
			)
			if shortestState == nil || len(currentState.Previous) < len(shortestState.Previous) {
				shortestState = currentState
			}
		}

		// Mark the current state as visited
		visited[currentState.Hash()] = true

		// Get the next possible states
		nextStates := currentState.GetNextPossibleStates()

		// Add the next possible states to the queue if they have not been visited
		numNew := 0
		for _, nextState := range nextStates {
			hash := nextState.Hash()
			if !visited[hash] && !willVisit[hash] {
				numNew++
				// fmt.Printf("\t\t[%d] %s\n", numNew, hash)
				willVisit[hash] = true
				queue = append(queue, nextState)
			}
		}

		// Remove the current state from the queue
		queue = queue[1:]

		iteration++
	}

	return shortestState, iteration
}
