package daily_challenge

import "testing"

func Test_MissingNumber(t *testing.T) {
	sourceCase := [][]int{
		{3, 0, 1},
		{9, 6, 4, 2, 3, 5, 7, 0, 1},
		{1, 2, 3},
	}
	expectCase := []int{
		2,
		8,
		0,
	}

	for i, source := range sourceCase {
		expect := expectCase[i]

		target := missingNumber(source)
		if target != expect {
			t.Error("Translate Test_MissingNumber Fail", expect, target)
		}
	}

	t.Log("Translate Test_MissingNumber Success")
}
