package weekly_contest

import "testing"

func Test_CarPooling(t *testing.T) {
	tripCase := [][][]int{
		[][]int{{2, 1, 5}, {3, 3, 7}},
		[][]int{{2, 1, 5}, {3, 3, 7}},
		[][]int{{2, 1, 5}, {3, 5, 7}},
		[][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}},
		[][]int{{7, 5, 6}, {6, 7, 8}, {10, 1, 6}},
	}
	capacityCase := []int{
		4, 5, 3, 11, 16,
	}
	expectList := []bool{
		false, true, true, true, false,
	}

	for i, source := range tripCase {
		expect := expectList[i]
		capacity := capacityCase[i]

		target := carPooling(source, capacity)
		if target != expect {
			t.Error("Run Test_CarPooling Fail", i, expect, target)
		}

		targetFast := carPoolingFast(source, capacity)
		if targetFast != expect {
			t.Error("Run Test_CarPoolingFast Fail", i, expect, targetFast)
		}
	}

	t.Log("Run Test_CarPooling Success")
}
