package weekly_contest_93

import (
	"testing"
)

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
