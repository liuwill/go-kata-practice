package weekly_contest

import "testing"

func Test_OddCells(t *testing.T) {
	nmCase := [][]int{
		{2, 3},
		{2, 2},
	}
	indiceCase := [][][]int{
		{{0, 1}, {1, 1}},
		{{1, 1}, {0, 0}},
	}
	expectList := []int{
		6,
		0,
	}

	for i, indices := range indiceCase {
		expect := expectList[i]
		nmPair := nmCase[i]

		target := oddCells(nmPair[0], nmPair[1], indices)
		if target != expect {
			t.Error("Run Test_OddCells Fail", expect, target)
		}
	}

	t.Log("Run Test_OddCells Success")
}
