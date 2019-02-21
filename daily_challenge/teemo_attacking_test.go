package daily_challenge

import "testing"

func Test_FindPoisonedDuration(t *testing.T) {
	sourceCase := [][]int{
		{1, 4},
		{1, 2},
	}
	attackCase := []int{
		2,
		2,
	}
	expectCase := []int{
		4,
		3,
	}

	for i, source := range sourceCase {
		expect := expectCase[i]
		attack := attackCase[i]

		target := findPoisonedDuration(source, attack)
		if target != expect {
			t.Error("Translate Test_FindPoisonedDuration Fail", expect, target)
		}
	}

	t.Log("Translate Test_FindPoisonedDuration Success")
}
