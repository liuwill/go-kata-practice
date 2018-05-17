package frog

import (
	"testing"
)

func Test_CanCross(t *testing.T) {
	source := []int{0, 1, 3, 5, 6, 8, 12, 17}
	target := canCross(source)
	expect := true

	if expect != target {
		t.Error("Translate Test_CanCross Fail")
	}
	t.Log("Translate Test_CanCross Success")
}

func Test_CanCrossFail(t *testing.T) {
	source := []int{0, 1, 2, 3, 4, 8, 9, 11}
	target := canCross(source)
	expect := false

	if expect != target {
		t.Error("Translate Test_CanCrossFail Fail", target)
	}
	t.Log("Translate Test_CanCrossFail more Success")
}
