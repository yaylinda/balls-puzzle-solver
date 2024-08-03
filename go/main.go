package main

import (
	"balls_puzzle_solver/pkg/puzzles"
	"balls_puzzle_solver/pkg/solver"
)

func main() {
	solver.Solve(puzzles.PUZZLE_COMPLETE)
	solver.Solve(puzzles.PUZZLE_EASY_1)
	solver.Solve(puzzles.PUZZLE_EASY_2)
}
