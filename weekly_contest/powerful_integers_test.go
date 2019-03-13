package weekly_contest

import (
	"sort"
	"testing"
)

func compareIntListWithoutOrder(source []int, target []int) bool {
	if len(source) != len(target) {
		return false
	}
	sort.Ints(source)
	sort.Ints(target)

	for i, letter := range source {
		if target[i] != letter {
			return false
		}
	}
	return true
}

func Test_PowerfulIntegers(t *testing.T) {
	sourceCase := [][]int{
		{2, 3, 10},
		{3, 5, 15},
	}
	expectList := [][]int{
		{2, 3, 4, 5, 7, 9, 10},
		{2, 4, 6, 8, 10, 14},
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := powerfulIntegers(source[0], source[1], source[2])
		if !compareIntListWithoutOrder(target, expect) {
			t.Error("Run Test_PowerfulIntegers Fail", expect, target)
		}
	}

	t.Log("Run Test_PowerfulIntegers Success")
}
