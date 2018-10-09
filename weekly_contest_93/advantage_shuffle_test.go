package weekly_contest_93

import (
	"testing"
)

func compareList(source []int, target []int) bool {
	if len(source) != len(target) {
		return false
	}

	for index := 0; index < len(source); index++ {
		if source[index] != target[index] {
			return false
		}
	}

	return true
}

func Test_AdvantageCount(t *testing.T) {
	sourceA := []int{12, 24, 8, 32}
	sourceB := []int{13, 25, 32, 11}
	target := advantageCount(sourceA, sourceB)
	expect := []int{24, 32, 8, 12}

	if !compareList(expect, target) {
		t.Error("Translate Test_AdvantageCount Fail", target)
	}
	t.Log("Translate Test_AdvantageCount Success")
}

func Test_AdvantageCountSimple(t *testing.T) {
	sourceA := []int{2, 7, 11, 15}
	sourceB := []int{1, 10, 4, 11}
	target := advantageCount(sourceA, sourceB)
	expect := []int{2, 11, 7, 15}

	if !compareList(expect, target) {
		t.Error("Translate Test_AdvantageCountSimple Fail", target)
	}
	t.Log("Translate Test_AdvantageCountSimple Success")
}

func Test_AdvantageCountFail(t *testing.T) {
	sourceA := []int{2, 0, 4, 1, 2}
	sourceB := []int{1, 3, 0, 0, 2}
	target := advantageCount(sourceA, sourceB)
	expect := []int{2, 0, 1, 2, 4}

	if !compareList(expect, target) {
		t.Error("Translate Test_AdvantageCountFail Fail", target)
	}
	t.Log("Translate Test_AdvantageCountFail Success")
}
