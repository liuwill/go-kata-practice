package rotate_image

import (
	"testing"
)

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

	for i := 0; i < len(source); i++ {
		for j := 0; j < len(source[i]); j++ {
			if source[i][j] != expect[i][j] {
				t.Error("Translate rotate Fail", i, j, source[i][j], expect[i][j])
				// break
			}
		}
	}
	t.Log("Translate Test_Rotate Success")
}
