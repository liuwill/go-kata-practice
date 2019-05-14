package weekly_contest

import (
	"testing"
)

func Test_GardenNoAdj(t *testing.T) {
	gardenCase := []int{
		3,
		4,
		4,
	}
	sourceCase := [][][]int{
		[][]int{
			[]int{1, 2},
			[]int{2, 3},
			[]int{3, 1},
		},
		[][]int{
			[]int{1, 2},
			[]int{3, 4},
		},
		[][]int{
			[]int{1, 2}, []int{2, 3}, []int{3, 4},
			[]int{4, 1}, []int{1, 3}, []int{2, 4},
		},
	}
	expectList := [][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 1, 2},
		[]int{1, 2, 3, 4},
	}

	for i, source := range sourceCase {
		expect := expectList[i]
		garden := gardenCase[i]

		target := gardenNoAdj(garden, source)
		if !compareList(target, expect) {
			t.Error("Run Test_GardenNoAdj Fail", expect, target)
		}
	}

	t.Log("Run Test_GardenNoAdj Success")
}
