package models

import (
	"strings"
)

func (b *Board) String() string {
	var builder strings.Builder

	for _, tower := range b.Towers {
		builder.WriteString("[")
		for i, ball := range tower.Balls {
			if i > 0 {
				builder.WriteString(",")
			}
			if ball == nil {
				builder.WriteString(" ")
			} else {
				builder.WriteString(ball.Color)
			}
		}
		builder.WriteString("]")
	}

	return builder.String()
}

func (bs *BoardState) isSolved(expectedEmpty int) bool {
	var numEmptyTowers int
	var numCompleteTowers int

	for _, tower := range bs.Board {
		if tower.isEmpty() {
			numEmptyTowers++
			continue
		}
		if tower.isComplete() {
			numCompleteTowers++
			continue
		}
	}

	return numEmptyTowers == expectedEmpty && numCompleteTowers == len(bs.Board)-expectedEmpty
}
