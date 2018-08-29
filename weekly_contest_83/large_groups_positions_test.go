package weekly_contest_83

import "testing"

func compareList(source []int, target []int) bool {
	if len(source) != len(target) {
		return false
	}

	for index := 0; index < len(source); index++ {
		if source[index] != target[index] {
			return false
		}
	}

	return true
}

func deepCompare(expect [][]int, target [][]int) bool {
	if len(expect) != len(target) {
		return false
	}
	for i := 0; i < len(expect); i++ {
		if !compareList(expect[i], target[i]) {
			return false
		}
	}
	return true
}

func Test_LargeGroupPositions(t *testing.T) {
	source := "abbxxxxzzy"
	target := largeGroupPositions(source)
	expect := [][]int{
		{3, 6},
	}

	if target == nil || !deepCompare(expect, target) {
		t.Error("Translate Test_LargeGroupPositions Fail", target)
	}
	t.Log("Translate Test_LargeGroupPositions Success")
}
