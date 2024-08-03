package main

import (
	"balls_puzzle_solver/pkg/models"
	"balls_puzzle_solver/pkg/puzzles"
)

func main() {
	// Load the puzzle
	towers := puzzles.CreateTowers(puzzles.PUZZLE_HARD)

	// Create a queue to keep track of the possible moves
	var queue []*models.BoardState

	// Populate the queue with all possible initial moves
	for i, t1 := range towers {
		for j, t2 := range towers {
			if towers[i].Index == towers[j].Index {
				continue
			}

			boardState := &models.BoardState{
				Board: towers,
				Move: models.Move{
					From: t1,
					To:   t2,
				},
				Previous: []*models.BoardState{},
			}

			queue = append(queue, boardState)
		}
	}

	// Keep track of the visited states
	visited := make(map[string]bool)

	// Iterate over the queue
	for len(queue) > 0 {
		// Get the first element of the queue
		currentState := queue[0]
		queue = queue[1:]

		// Check if the current state is the final state
		if currentState.IsFinalState() {
			// Print the path to the solution
			currentState.PrintPath()
			break
		}

		// Check if the current state has been visited
		if visited[currentState.Hash()] {
			continue
		}

		// Mark the current state as visited
		visited[currentState.Hash()] = true

		// Get the next possible moves
		nextStates := currentState.GetNextStates()

		// Add the next possible moves to the queue
		queue = append(queue, nextStates...)
	}
}
