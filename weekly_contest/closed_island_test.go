package weekly_contest

import "testing"

func Test_ClosedIsland(t *testing.T) {
	sourceList := [][][]int{
		{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}},
		{{0, 0, 1, 0, 0}, {0, 1, 0, 1, 0}, {0, 1, 1, 1, 0}},
		{
			{1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 0, 1},
			{1, 0, 1, 0, 1, 0, 1},
			{1, 0, 1, 1, 1, 0, 1},
			{1, 0, 0, 0, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1},
		},
	}
	expectList := []int{
		2,
		1,
		2,
	}

	for i, source := range sourceList {
		expect := expectList[i]
		target := closedIsland(source)

		if expect != target {
			t.Error("Run ClosedIsland Fail ", target)
		}
	}

	t.Log("Test_ClosedIsland Success")
}
