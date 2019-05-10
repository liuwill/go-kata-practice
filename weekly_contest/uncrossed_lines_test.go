package weekly_contest

import "testing"

func Test_MaxUncrossedLines(t *testing.T) {
	sourceCase := [][][]int{
		[][]int{
			{1, 4, 2},
			{1, 2, 4},
		},
	}
	expectList := []int{
		3,
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
