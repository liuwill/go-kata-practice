package daily_challenge

import "testing"

func Test_FindJudge(t *testing.T) {
	numSource := []int{
		2, 3, 3, 3, 4,
	}
	trustSource := [][][]int{
		{{1, 2}},
		{{1, 3}, {2, 3}},
		{{1, 3}, {2, 3}, {3, 1}},
		{{1, 2}, {2, 3}},
		{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}},
	}
	expectList := []int{
		2, 3, -1, -1, 3,
	}

	for i, source := range numSource {
		trustList := trustSource[i]
		expect := expectList[i]

		target := findJudge(source, trustList)
		if target != expect {
			t.Error("Run Test_FindJudge Fail", expect, target)
		}
	}

	t.Log("Run Test_FindJudge Success")
}
