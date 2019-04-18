package daily_challenge

import "testing"

func Test_Divide(t *testing.T) {
	sourceCase := [][]int{
		{10, 3},
		{7, -3},
		{-2147483648, -1},
		{-2147483648, 2},
	}
	expectList := []int{
		3,
		-2,
		2147483647,
		-1073741824,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := divide(source[0], source[1])
		if target != expect {
			t.Error("Run Test_divide Fail", source, expect, target)
		}
	}

	t.Log("Run Test_divide Success")
}
