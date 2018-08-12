package weekly_contest_93

import (
	"testing"
)

func Test_SortMark(t *testing.T) {
	source := []int{12, 24, 8, 32}
	target := sortMark(source)
	expect := []int{8, 12, 24, 32}

	realTarget := []int{}
	for _, item := range target {
		realTarget = append(realTarget[:], item.Number)
	}

	if !compareList(expect, realTarget) {
		t.Error("Translate Test_SortMark Fail", realTarget)
	}
	t.Log("Translate Test_SortMark Success")
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
