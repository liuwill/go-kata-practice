package daily_challenge

import "testing"

func Test_Divide(t *testing.T) {
	sourceCase := [][]int{
		{10, 3},
		{7, -3},
	}
	expectList := []int{
		3,
		-2,
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
