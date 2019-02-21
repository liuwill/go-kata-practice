package daily_challenge

import "testing"

func Test_CanPlaceFlowers(t *testing.T) {
	sourceCase := [][]int{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 1, 0, 0},
	}
	attackCase := []int{
		1,
		2,
		2,
		1,
		1,
	}
	expectCase := []bool{
		true,
		false,
		false,
		true,
		true,
	}

	for i, source := range sourceCase {
		expect := expectCase[i]
		attack := attackCase[i]

		target := canPlaceFlowers(source, attack)
		if target != expect {
			t.Error("Translate Test_CanPlaceFlowers Fail", expect, target)
		}
	}

	t.Log("Translate Test_CanPlaceFlowers Success")
}
