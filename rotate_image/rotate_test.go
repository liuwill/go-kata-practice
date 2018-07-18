package rotate_image

import (
	"testing"
)

func isArrayMatch(target [][]int, expect [][]int) bool {
	for i := 0; i < len(target); i++ {
		for j := 0; j < len(target[i]); j++ {
			if target[i][j] != expect[i][j] {
				return false
			}
		}
	}
	return true
}

func Test_Rotate(t *testing.T) {
	source := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(source)
	expect := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	if !isArrayMatch(source, expect) {
		t.Error("Translate rotate Fail", source)
	}
	t.Log("Translate Test_Rotate Success")
}

func Test_RotateMore(t *testing.T) {
	source := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(source)
	expect := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}

	if !isArrayMatch(source, expect) {
		t.Error("Translate rotate more Fail", source)
	}
	t.Log("Translate Test_RotateMore Success")
}
