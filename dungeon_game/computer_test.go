package dungeon_game

import (
	"testing"
)

func Test_ComputerMiniHp(t *testing.T) {
	source := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	target := calculateComputerMinimumHP(source)
	expect := 7

	if expect != target {
		t.Error("Translate computer minimum hp Fail", target)
	}
	t.Log("Translate computer Test_ComputerMiniHp Success")
}

func Test_ComputerMiniHpZero(t *testing.T) {
	source := [][]int{{0, 0}}
	target := calculateComputerMinimumHP(source)
	expect := 1

	if expect != target {
		t.Error("Translate computer minimum hp zero Fail", target)
	}
	t.Log("Translate computer Test_ComputerMiniHpZero Success")
}

func Test_ComputerMiniHpError(t *testing.T) {
	source := [][]int{{2}, {1}}
	target := calculateComputerMinimumHP(source)
	expect := 1

	if expect != target {
		t.Error("Translate computer minimum hp error Fail", target)
	}
	t.Log("Translate computer Test_ComputerMiniHpError Success")
}
