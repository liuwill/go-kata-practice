package weekly_contest

import "testing"

func Test_CarPooling(t *testing.T) {
	tripCase := [][][]int{
		[][]int{{2, 1, 5}, {3, 3, 7}},
		[][]int{{2, 1, 5}, {3, 3, 7}},
		[][]int{{2, 1, 5}, {3, 5, 7}},
		[][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}},
	}
	capacityCase := []int{
		4, 5, 3, 11,
	}
	expectList := []bool{
		false, true, true, true,
	}

	for i, source := range tripCase {
		expect := expectList[i]
		capacity := capacityCase[i]

		target := carPooling(source, capacity)
		if target != expect {
			t.Error("Run Test_CarPooling Fail", i, expect, target)
		}
	}

	t.Log("Run Test_CarPooling Success")
}
