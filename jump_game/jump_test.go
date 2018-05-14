package jump_game

import (
	"testing"
)

func Test_Jump(t *testing.T) {
	source := []int{2, 3, 1, 1, 4}
	target := jump(source)
	expect := 2

	if expect != target {
		t.Error("Translate jump game II Fail ", target)
	}
	t.Log("Translate Test_Jump Success")
}

func Test_JumpMore(t *testing.T) {
	source := []int{2, 4, 2, 1, 5, 3, 2, 6, 3}
	target := jump(source)
	expect := 3

	if expect != target {
		t.Error("Translate jump game II more Fail ", target)
	}
	t.Log("Translate Test_JumpMore more Success")
}

func Test_JumpLess(t *testing.T) {
	source := []int{2, 1}
	target := jump(source)
	expect := 1

	if expect != target {
		t.Error("Translate jump game II less Fail ", target)
	}
	t.Log("Translate Test_JumpLess Success")
}

func Test_JumpGameOne(t *testing.T) {
	source := []int{2}
	target := jump(source)
	expect := 0

	if expect != target {
		t.Error("Translate jump game II one Fail ")
	}
	t.Log("Translate Test_JumpGameOne Success")
}

func Test_JumpGameEmpty(t *testing.T) {
	source := []int{}
	target := jump(source)
	expect := 0

	if expect != target {
		t.Error("Translate jump game II empty Fail ")
	}
	t.Log("Translate Test_JumpGameEmpty Success")
}
