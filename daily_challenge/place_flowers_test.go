package daily_challenge

import "testing"

func copyArray(source []int) []int {
	target := make([]int, len(source))
	copy(target, source)

	return target
}

func Test_CanPlaceFlowers(t *testing.T) {
	sourceCase := [][]int{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 1, 0, 0},
		{1, 0, 0, 0, 1, 0, 0},
	}
	attackCase := []int{
		1,
		2,
		2,
		1,
		1,
		2,
	}
	expectCase := []bool{
		true,
		false,
		false,
		true,
		true,
		true,
	}

	for i, source := range sourceCase {
		expect := expectCase[i]
		attack := attackCase[i]

		target := canPlaceFlowers(copyArray(source), attack)
		if target != expect {
			t.Error("Translate Test_CanPlaceFlowers Fail", expect, target)
		}

		targetSlow := canPlaceFlowersSlow(copyArray(source), attack)
		if targetSlow != expect {
			t.Error("Translate Test_CanPlaceFlowersSlow Fail", source, expect, targetSlow)
		}
	}

	t.Log("Translate Test_CanPlaceFlowers Success")
}
