package weekly_contest

import "testing"

func Test_ShiftGrid(t *testing.T) {
	gridCase := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}
	kCase := []int{1, 4, 9}
	expectList := [][][]int{
		{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}

	for i, grid := range gridCase {
		expect := expectList[i]

		target := shiftGrid(grid, kCase[i])
		if !comparePointMatch(target, expect) {
			t.Error("Run Test_ShiftGrid Fail", expect, target)
		}
	}

	t.Log("Run Test_ShiftGrid Success")
}
