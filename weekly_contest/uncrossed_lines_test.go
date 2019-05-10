package weekly_contest

import "testing"

func Test_MaxUncrossedLines(t *testing.T) {
	sourceCase := [][][]int{
		[][]int{
			{1, 4, 2},
			{1, 2, 4},
		},
		[][]int{
			{2, 5, 1, 2, 5},
			{10, 5, 2, 1, 5, 2},
		},
		[][]int{
			{1, 3, 7, 1, 7, 5},
			{1, 9, 2, 5, 1},
		},
	}
	expectList := []int{
		2,
		3,
		2,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := maxUncrossedLines(source[0], source[1])
		if target != expect {
			t.Error("Run Test_MaxUncrossedLines Fail", expect, target)
		}
	}

	t.Log("Run Test_MaxUncrossedLines Success")
}
