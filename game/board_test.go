package game

import "testing"

func TestBoardScore(t *testing.T) {
	board := NewBoard(4,
		Ship{Coord{0, 1}},
		Ship{Coord{2, 0}},
		Ship{Coord{2, 1}},
	)
	board.AddAttack(2, 1)

	if score := board.Score(); score != 1 {
		t.Error("the score expected to be 1 got ", score)
	}

	board.AddAttack(1, 1)

	if score := board.Score(); score != 1 {
		t.Error("the score expected to be 1 got ", score)
	}
	board.AddAttack(2, 0)

	if score := board.Score(); score != 2 {
		t.Error("the score expected to be 2 got ", score)
	}

}
