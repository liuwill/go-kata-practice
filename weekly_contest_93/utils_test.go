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
