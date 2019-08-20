package weekly_contest

import "testing"

func Test_MaxDistance(t *testing.T) {
	mapCase := [][][]int{
		{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	}
	expectList := []int{2, 4, -1}

	for i, source := range mapCase {
		expect := expectList[i]

		target := maxDistance(source)
		if target != expect {
			t.Error("Run Test_MaxDistance Fail", expect, target)
		}
	}

	t.Log("Run Test_MaxDistance Success")
}
