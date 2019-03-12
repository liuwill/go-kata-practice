package weekly_contest

import "testing"

func Test_OrangesRotting(t *testing.T) {
	sourceCase := [][][]int{
		{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
		{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
		{{0, 2}},
	}
	expectList := []int{
		4,
		-1,
		0,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := orangesRotting(source)
		if target != expect {
			t.Error("Run Test_OrangesRotting Fail", expect, target)
		}
	}

	t.Log("Run Test_OrangesRotting Success")
}
