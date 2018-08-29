package weekly_contest_83

import "testing"

func deepCompare(expect [][]int, target [][]int) bool {
	return false
}

func Test_LargeGroupPositions(t *testing.T) {
	source := "abbxxxxzzy"
	target := largeGroupPositions(source)
	expect := [][]int{
		{3, 6},
	}

	if !deepCompare(expect, target) {
		t.Error("Translate Test_LargeGroupPositions Fail", target)
	}
	t.Log("Translate Test_LargeGroupPositions Success")
}
