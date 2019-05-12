package weekly_contest

import (
	"testing"
)

func Test_NumMovesStones(t *testing.T) {
	sourceCase := [][]int{
		[]int{1, 2, 5},
		[]int{4, 3, 2},
		[]int{3, 5, 1},
	}
	expectList := [][]int{
		[]int{1, 2},
		[]int{0, 0},
		[]int{1, 2},
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := numMovesStones(source[0], source[1], source[2])
		if len(target) != 2 || (target[0] != expect[0] || target[1] != expect[1]) {
			t.Error("Run Test_NumMovesStones Fail", expect, target)
		}
	}

	t.Log("Run Test_NumMovesStones Success")
}
