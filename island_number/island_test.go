package island_number

import (
	"testing"
)

func Test_Island(t *testing.T) {
	source := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	target := numIslands(source)
	expect := 3

	if expect != target {
		t.Error("Translate island number Fail ", target)
	}
	t.Log("Translate Test_Island Success")
}
