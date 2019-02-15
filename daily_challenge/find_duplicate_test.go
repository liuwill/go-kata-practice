package daily_challenge

import "testing"

func Test_FindDuplicate(t *testing.T) {
	sourceCase := [][]int{
		{1, 3, 4, 2, 2},
		{3, 1, 3, 4, 2},
	}
	expectCase := []int{
		2,
		3,
	}

	for i, source := range sourceCase {
		expect := expectCase[i]

		target := findDuplicate(source)
		if target != expect {
			t.Error("Translate Test_FindDuplicate Fail", expect, target)
		}

		optimizeTarget := findDuplicateOptimize(source)
		if optimizeTarget != expect {
			t.Error("Translate Test_FindDuplicateOptimize Fail", expect, optimizeTarget)
		}
	}

	t.Log("Translate Test_FindDuplicate Success")
}
