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

func Test_LargeGroupPositionsBlank(t *testing.T) {
	source := "abc"
	target := largeGroupPositions(source)
	expect := [][]int{}

	if target == nil || !deepCompare(expect, target) {
		t.Error("Translate Test_LargeGroupPositionsBlank Fail", target)
	}
	t.Log("Translate Test_LargeGroupPositionsBlank Success")
}

func Test_LargeGroupPositionsMore(t *testing.T) {
	source := "abcdddeeeeaabbbcd"
	target := largeGroupPositions(source)
	expect := [][]int{
		{3, 5},
		{6, 9},
		{12, 14},
	}

	if target == nil || !deepCompare(expect, target) {
		t.Error("Translate Test_LargeGroupPositionsMore Fail", target)
	}
	t.Log("Translate Test_LargeGroupPositionsMore Success")
}

func Test_LargeGroupPositionsBorder(t *testing.T) {
	source := "abcdddeeeeaabbbcddd"
	target := largeGroupPositions(source)
	expect := [][]int{
		{3, 5},
		{6, 9},
		{12, 14},
		{16, 18},
	}

	if target == nil || !deepCompare(expect, target) {
		t.Error("Translate Test_LargeGroupPositionsBorder Fail", target)
	}
	t.Log("Translate Test_LargeGroupPositionsBorder Success")
}
