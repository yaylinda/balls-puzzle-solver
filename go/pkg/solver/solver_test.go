package solver

import (
	"testing"

	"balls_puzzle_solver/pkg/puzzles"
	"github.com/stretchr/testify/assert"
)

func TestSolveEasy(t *testing.T) {
	tests := []struct {
		name     string
		puzzle   [][]string
		expected int
	}{
		{"PUZZLE_COMPLETE", puzzles.PUZZLE_COMPLETE, 0},
		{"PUZZLE_EASY_1", puzzles.PUZZLE_EASY_1, 1},
		{"PUZZLE_EASY_2", puzzles.PUZZLE_EASY_2, 2},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				a := assert.New(t)
				s, iters := Solve(tt.puzzle, PrintOpts{})
				a.NotNil(s)
				a.Equal(tt.expected, len(s.Previous))
				t.Logf("Iterations: %d", iters)
			},
		)
	}
}
