package weekly_contest

import "testing"

func Test_CorpFlightBookings(t *testing.T) {
	rootCase := [][][]int{
		[][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}},
		[][]int{{1, 2, 10}, {2, 3, 20}, {2, 4, 25}},
	}
	sourceCase := []int{
		5,
		5,
	}
	expectList := [][]int{
		[]int{10, 55, 45, 25, 25},
		[]int{10, 55, 45, 25, 0},
	}

	for i, source := range sourceCase {
		root := rootCase[i]
		expect := expectList[i]

		target := corpFlightBookings(root, source)
		if !compareList(target, expect) {
			t.Error("Run Test_CorpFlightBookings Fail", expect, target, root)
		}
	}

	t.Log("Run Test_CorpFlightBookings Success")
}
