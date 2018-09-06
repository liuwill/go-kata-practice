package daily_challenge

import (
	"testing"
)

func Test_CompareList(t *testing.T) {
	source1 := []int{12, 24}
	source2 := []int{12, 24}
	target := compareList(source1, source2)
	expect := true

	if expect != target {
		t.Error("Translate Test_CompareList Fail", target)
	}
	t.Log("Translate Test_CompareList Success")
}

func Test_CompareListLength(t *testing.T) {
	source1 := []int{12, 24, 8, 32}
	source2 := []int{12, 24}
	target := compareList(source1, source2)
	expect := false

	if expect != target {
		t.Error("Translate Test_CompareListLength Fail", target)
	}
	t.Log("Translate Test_CompareListLength Success")
}

func Test_CompareListFail(t *testing.T) {
	source1 := []int{12, 24, 8, 32}
	source2 := []int{12, 8, 12, 44}
	target := compareList(source1, source2)
	expect := false

	if expect != target {
		t.Error("Translate Test_CompareListFail Fail", target)
	}
	t.Log("Translate Test_CompareListFail Success")
}
