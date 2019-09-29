package weekly_contest

import "testing"

func Test_UniqueOccurrences(t *testing.T) {
	sourceCase := [][]int{
		{1, 2, 2, 1, 1, 3},
		{1, 2},
		{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
	}
	expectList := []bool{
		true,
		false,
		true,
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := uniqueOccurrences(source)
		if target != expect {
			t.Error("Run Test_UniqueOccurrences Fail", expect, target, i)
		}
	}

	t.Log("Run Test_UniqueOccurrences Success")
}
