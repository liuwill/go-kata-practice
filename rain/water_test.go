package rain

import (
	"testing"
)

func Test_Trap(t *testing.T) {
	source := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	target := trap(source)
	expect := 6

	if expect != target {
		t.Error("Translate trap rain water Fail ")
	}
	t.Log("Translate trap rain water Success")
}
