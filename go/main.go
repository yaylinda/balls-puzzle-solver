package main

import (
	"fmt"

	"balls_puzzle_solver/pkg/models"
	"balls_puzzle_solver/pkg/puzzles"
)

func main() {
	// Create a board from the puzzle
	initialBoard := models.CreateBoardFromPuzzle(puzzles.PUZZLE_EASY_1)

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

	// Iterate over the queue
	for len(queue) > 0 {
		fmt.Printf("iter: %d, queue: %d\n", iteration, len(queue))

		// Get the first element of the queue
		currentState := queue[0]

		// Check if the current state is the final state
		if currentState.Board.IsSolved(2) {
			fmt.Println(currentState.PrintSolution())
			break
		}

		// Mark the current state as visited
		visited[currentState.String()] = true

		// Get the next possible states
		nextStates := currentState.GetNextPossibleStates()

		// Add the next possible states to the queue if they have not been visited
		numNew := 0
		for _, nextState := range nextStates {
			hash := nextState.String()
			if !visited[hash] && !willVisit[hash] {
				numNew++
				// fmt.Printf("\t\t[%d] %s\n", numNew, hash)
				willVisit[hash] = true
				queue = append(queue, nextState)
			}
		}

		fmt.Printf("\tadded %d / %d next possible states\n", numNew, len(nextStates))

		// Remove the current state from the queue
		queue = queue[1:]

		iteration++
	}
}
