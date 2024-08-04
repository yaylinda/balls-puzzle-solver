package main

import (
	"log"

	"balls_puzzle_solver/pkg/puzzles"
	"balls_puzzle_solver/pkg/solver"
	"balls_puzzle_solver/pkg/visualizer"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	s, _ := solver.Solve(puzzles.PUZZLE_HARD, solver.PrintOpts{})

	ebiten.SetWindowSize(visualizer.ScreenWidth, visualizer.ScreenHeight)
	ebiten.SetWindowTitle("Color Grid Visualization")

	if err := ebiten.RunGame(
		&visualizer.Visualizer{
			CurrentIndex: 0,
			States:       s.GetSolvedPath(),
		},
	); err != nil {
		log.Fatal(err)
	}
}
