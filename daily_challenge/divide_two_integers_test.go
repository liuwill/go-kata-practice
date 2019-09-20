package daily_challenge

import (
	"math"
	"testing"
)

func Test_Divide(t *testing.T) {
	max := int(math.Pow(2, 31) - 1)
	min := int(math.Pow(2, 31)) * -1

	sourceCase := [][]int{
		{10, 3},
		{7, -3},
		{-2147483648, -1},
		{-2147483648, 2},
		{min - 1, 1},
	}
	expectList := []int{
		3,
		-2,
		max,
		-1073741824,
		min,
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
