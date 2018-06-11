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
	source := [][]int{{0, -5}, {0, 0}}
	target := calculateMinimumHP(source)
	expect := 1

	if expect != target {
		t.Error("Translate minimum hp fail Fail", target)
	}
	t.Log("Translate Test_MiniHpFails Success")
}
