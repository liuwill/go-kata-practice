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

func Test_TrapLess(t *testing.T) {
	source := []int{0, 1}
	target := trap(source)
	expect := 0

	if expect != target {
		t.Error("Translate trap rain water less Fail ")
	}
	t.Log("Translate Test_TrapLess Success")
}

func Test_TrapHighFirst(t *testing.T) {
	source := []int{4, 1, 2, 1, 2}
	target := trap(source)
	expect := 2

	if expect != target {
		t.Error("Translate trap rain water HighFirst Fail ")
	}
	t.Log("Translate Test_TrapHighFirst Success")
}

func Test_TrapHighLast(t *testing.T) {
	source := []int{2, 1, 4}
	target := trap(source)
	expect := 1

	if expect != target {
		t.Error("Translate trap rain water HighLast Fail ")
	}
	t.Log("Translate Test_TrapHighLast Success")
}
