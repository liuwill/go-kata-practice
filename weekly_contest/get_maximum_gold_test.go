package weekly_contest

import "testing"

func Test_GetMaximumGold(t *testing.T) {
	sourceCase := [][][]int{
		{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}},
		{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}},
		{
			{1, 0, 7, 0, 0, 0},
			{2, 0, 6, 0, 1, 0},
			{3, 5, 6, 7, 4, 2},
			{4, 0, 0, 0, 0, 0},
			{3, 0, 0, 20, 20, 20},
		},
	}
	expectList := []int{
		24,
		28,
		60,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := getMaximumGold(source)
		if target != expect {
			t.Error("Run Test_GetMaximumGold Fail", expect, target)
		}
	}

	t.Log("Run Test_GetMaximumGold Success")
}
