package daily_challenge

import (
	"testing"
)

func Test_TwoSum(t *testing.T) {
	source := []int{2, 7, 11, 15}
	sum := 9
	target := twoSum(source, sum)
	expect := []int{0, 1}

	if !compareList(expect, target) {
		t.Error("Translate Test_TwoSum Fail", target)
	}
	t.Log("Translate Test_TwoSum Success")
}

func Test_TwoSumEmpty(t *testing.T) {
	source := []int{2, 7, 11, 15}
	sum := 3
	target := twoSum(source, sum)
	expect := []int{0, 0}

	if !compareList(expect, target) {
		t.Error("Translate Test_TwoSum Fail", target)
	}
	t.Log("Translate Test_TwoSum Success")
}
