package main

import (
	"balls_puzzle_solver/pkg/puzzles"
	"balls_puzzle_solver/pkg/solver"
	"balls_puzzle_solver/pkg/utils"
)

func main() {
	solver.Solve(puzzles.PUZZLE_HARD, solver.PrintOpts{Mod: utils.Ptr(10_000)})
}
