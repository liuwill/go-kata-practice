package weekly_contest_79

import (
	"math"
	"testing"
)

func equals(expect float64, target float64) bool {
	return math.Round(expect*1000000) == math.Round(target*1000000)
}

func Test_LargestTriangleArea(t *testing.T) {
	source := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
	target := largestTriangleArea(source)
	expect := 2.0

	if !equals(expect, target) {
		t.Error("Translate Test_LargestTriangleArea Fail", target)
	}
	t.Log("Translate Test_LargestTriangleArea Success")
}

func Test_IsTriangle(t *testing.T) {
	source := [][]int{{0, 0}, {0, 1}, {0, 2}}
	target := isTriangle(source[0], source[1], source[2])
	expect := false

	if expect != target {
		t.Error("Translate Test_IsTriangle Fail", target)
	}
	t.Log("Translate Test_IsTriangle Success")
}

func Test_IsTriangleSame(t *testing.T) {
	source := [][]int{{0, 2}, {1, 1}, {0, 2}}
	target := isTriangle(source[0], source[1], source[2])
	expect := false

	if expect != target {
		t.Error("Translate Test_IsTriangleSame Fail", target)
	}
	t.Log("Translate Test_IsTriangleSame Success")
}
