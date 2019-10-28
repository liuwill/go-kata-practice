package weekly_contest

import "testing"

func Test_CheckStraightLine(t *testing.T) {
	sourceCase := [][][]int{
		{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}},
		{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}},
		{{1, 1}, {1, 2}, {1, 4}, {1, 5}},
		{{1, 2}, {3, 2}, {6, 2}, {9, 2}},
		{{1, 2}, {3, 2}, {6, 3}, {9, 2}},
	}
	expectList := []bool{
		true, false, true, true, false,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := checkStraightLine(source)
		if target != expect {
			t.Error("Run Test_CheckStraightLine Fail", i, expect, target)
		}
	}

	t.Log("Run Test_CheckStraightLine Success")
}
