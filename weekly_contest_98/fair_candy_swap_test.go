package weekly_contest_98

import "testing"

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
func Test_FairCandySwap(t *testing.T) {
	sourceA := []int{1, 1}
	sourceB := []int{2, 2}
	target := fairCandySwap(sourceA, sourceB)
	expect := []int{1, 2}

	if !compareList(expect, target) {
		t.Error("Translate Test_FairCandySwap Fail", target)
	}
	t.Log("Translate Test_FairCandySwap Success")
}

func Test_FairCandySwapSecond(t *testing.T) {
	sourceA := []int{1, 2, 5}
	sourceB := []int{2, 4}
	target := fairCandySwap(sourceA, sourceB)
	expect := []int{5, 4}

	if !compareList(expect, target) {
		t.Error("Translate Test_FairCandySwapSecond Fail", target)
	}
	t.Log("Translate Test_FairCandySwapSecond Success")
}

func Test_FairCandySwapLess(t *testing.T) {
	sourceA := []int{2}
	sourceB := []int{1, 3}
	target := fairCandySwap(sourceA, sourceB)
	expect := []int{2, 3}

	if !compareList(expect, target) {
		t.Error("Translate Test_FairCandySwapLess Fail", target)
	}
	t.Log("Translate Test_FairCandySwapLess Success")
}
