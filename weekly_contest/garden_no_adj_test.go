package weekly_contest

import (
	"testing"
)

func Test_GardenNoAdj(t *testing.T) {
	gardenCase := []int{
		3,
	}
	sourceCase := [][][]int{
		[][]int{
			[]int{1, 2},
			[]int{2, 3},
			[]int{3, 1},
		},
	}
	expectList := [][]int{
		[]int{1, 2, 3},
	}

	for i, source := range sourceCase {
		expect := expectList[i]
		garden := gardenCase[i]

		target := gardenNoAdj(garden, source)
		if compareList(target, expect) {
			t.Error("Run Test_GardenNoAdj Fail", expect, target)
		}
	}

	t.Log("Run Test_GardenNoAdj Success")
}
