package weekly_contest_79

import (
	"testing"
)

func Test_LargestTriangleArea(t *testing.T) {
	source := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	target := largestTriangleArea(source)
	expect := 2.0

	if expect != target {
		t.Error("Translate Test_LargestTriangleArea Fail", target)
	}
	t.Log("Translate Test_LargestTriangleArea Success")
}
