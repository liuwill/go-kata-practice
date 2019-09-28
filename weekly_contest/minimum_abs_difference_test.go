package weekly_contest

import (
	"testing"
)

func compareInnerArray(expect [][]int, target [][]int) bool {
	if len(expect) != len(target) {
		return false
	}

	for i, val := range expect {
		tar := target[i]
		if len(tar) != len(val) {
			return false
		}

		for j, v := range val {
			if v != tar[j] {
				return false
			}
		}
	}

	return true
}

func Test_MinimumAbsDifference(t *testing.T) {
	sourceCase := [][]int{
		{4, 2, 1, 3},
		{1, 3, 6, 10, 15},
		{3, 8, -10, 23, 19, -4, -14, 27},
	}
	expectList := [][][]int{
		{{1, 2}, {2, 3}, {3, 4}},
		{{1, 3}},
		{{-14, -10}, {19, 23}, {23, 27}},
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := minimumAbsDifference(source)
		if !compareInnerArray(target, expect) {
			t.Error("Run Test_MinimumAbsDifference Fail", expect, target)
		}
	}

	t.Log("Run Test_MinimumAbsDifference Success")
}
