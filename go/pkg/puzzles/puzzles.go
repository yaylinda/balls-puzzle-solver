package puzzles

import (
	"fmt"

	"balls_puzzle_solver/pkg/models"
)

var PUZZLE_HARD = [][]string{
	{"brown", "green", "brown", "light green"},
	{"light green", "gray", "light blue", "green"},
	{"red", "green", "light green", "yellow"},
	{"pink", "red", "yellow", "purple"},
	{"light green", "pink", "red", "green"},
	{"gray", "blue", "pink", "gray"},
	{"orange", "blue", "yellow", "dark green"},
	{"purple", "orange", "orange", "dark green"},
	{"light blue", "blue", "orange", "red"},
	{"purple", "dark green", "dark green", "light blue"},
	{"yellow", "pink", "purple", "blue"},
	{"light blue", "brown", "brown", "gray"},
	{"", "", "", ""},
	{"", "", "", ""},
}

// CreateTowers translates a "puzzle" into an array of Towers
func CreateTowers(balls [][]string) []*models.Tower {
	colorCount := make(map[string]int)
	towers := make([]*models.Tower, len(balls))

	for i, towerBalls := range balls {
		tower := &models.Tower{Index: i, Balls: make([]*models.Ball, len(towerBalls))}
		for j, color := range towerBalls {
			if color != "" {
				colorCount[color]++
				id := fmt.Sprintf("%s_%d", color, colorCount[color])
				tower.Balls[j] = &models.Ball{ID: id, Color: color}
			}
		}
		towers[i] = tower
	}

	return towers
}
