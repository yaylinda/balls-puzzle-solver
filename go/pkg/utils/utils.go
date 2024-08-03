package utils

import (
	"fmt"

	"balls_puzzle_solver/pkg/models"
)

// CreateBoardFromPuzzle creates a Board from a puzzle
func CreateBoardFromPuzzle(puzzle [][]string) *models.Board {
	colorCount := make(map[string]int)
	towers := make([]*models.Tower, len(puzzle))

	for i, towerBalls := range puzzle {
		tower := &models.Tower{
			Index: i,
			Balls: make([]*models.Ball, len(towerBalls)),
		}
		for j, color := range towerBalls {
			if color != "" {
				colorCount[color]++
				id := fmt.Sprintf("%s_%d", color, colorCount[color])
				tower.Balls[j] = &models.Ball{ID: id, Color: color}
			}
		}
		towers[i] = tower
	}

	return &models.Board{Towers: towers}
}

func Ptr[K any](x K) *K {
	return &x
}
