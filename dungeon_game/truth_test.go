package dungeon_game

import (
	"testing"
)

func Test_MiniHp(t *testing.T) {
	source := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	target := calculateMinimumHP(source)
	expect := 7

	if expect != target {
		t.Error("Translate minimum hp Fail", target)
	}
	t.Log("Translate Test_MiniHp Success")
}

func Test_MiniHpFails(t *testing.T) {
	source := [][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}
	target := calculateMinimumHP(source)
	expect := 3

	if expect != target {
		t.Error("Translate minimum hp fail Fail", target)
	}
	t.Log("Translate Test_MiniHpFails Success")
}

func Test_MiniHpZero(t *testing.T) {
	source := [][]int{{0, 0}}
	target := calculateMinimumHP(source)
	expect := 1

	if expect != target {
		t.Error("Translate minimum hp zero Fail", target)
	}
	t.Log("Translate Test_MiniHpZero Success")
}

func Test_MiniHpError(t *testing.T) {
	source := [][]int{{2}, {1}}
	target := calculateMinimumHP(source)
	expect := 1

	if expect != target {
		t.Error("Translate minimum hp error Fail", target)
	}
	t.Log("Translate Test_MiniHpError Success")
}
