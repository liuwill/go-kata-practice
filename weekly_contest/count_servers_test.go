package weekly_contest

import "testing"

func Test_CountServers(t *testing.T) {
	gridCase := [][][]int{
		{{1, 0}, {0, 1}},
		{{1, 0}, {1, 1}},
		{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
		{{1, 0, 0, 1, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}},
	}
	expectList := []int{0, 3, 4, 3}

	for i, grid := range gridCase {
		expect := expectList[i]

		target := countServers(grid)
		if target != expect {
			t.Error("Run Test_CountServers Fail", expect, target)
		}
	}

	t.Log("Run Test_CountServers Success")
}
