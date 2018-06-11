package dungeon_game

import (
	"testing"
)

func Test_KnightMiniHp(t *testing.T) {
	source := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	target := calculateKnightMinimumHP(source)
	expect := 7

	if expect != target {
		t.Error("Translate knight minimum hp Fail", target)
	}
	t.Log("Translate Test_MiniHp knight Success")
}

func Test_KnightMiniHpFails(t *testing.T) {
	source := [][]int{{0, -5}, {0, 0}}
	target := calculateKnightMinimumHP(source)
	expect := 1

	if expect != target {
		t.Error("Translate knight minimum hp fail Fail", target)
	}
	t.Log("Translate Test_MiniHpFails knight Success")
}

func Test_KnightMiniHpError(t *testing.T) {
	source := [][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}
	target := calculateKnightMinimumHP(source)
	expect := 3

	if expect != target {
		t.Error("Translate knight minimum hp error Fail", target)
	}
	t.Log("Translate Test_KnightMiniHpError Success")
}
