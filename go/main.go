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

	viz := &visualizer.Visualizer{
		CurrentIndex:        0,
		States:              s.GetSolvedPath(),
		UnitSize:            50,
		ButtonWidth:         50 * 3,
		ScreenWidth:         50 * len(puzzles.PUZZLE_HARD) / 2,
		ScreenHeight:        50 * (len(puzzles.PUZZLE_HARD[0])*2 + 3),
		SecondRowTowerIndex: len(puzzles.PUZZLE_HARD) / 2,
		SecondRowOffset:     float32(len(puzzles.PUZZLE_HARD[0])) + 0.5,
	}

	ebiten.SetWindowSize(viz.ScreenWidth, viz.ScreenHeight)
	ebiten.SetWindowTitle("Color Grid Visualization")

	if err := ebiten.RunGame(viz); err != nil {
		log.Fatal(err)
	}
}
