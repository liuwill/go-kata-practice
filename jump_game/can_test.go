package jump_game

import (
	"testing"
)

func Test_CanJumpSuccess(t *testing.T) {
	source := []int{2, 3, 1, 1, 4}
	target := canJump(source)
	expect := true

	if expect != target {
		t.Error("Translate Jump Game CanSuccess Fail ", target)
	}
	t.Log("Translate Test_CanJumpSuccess Success")
}

func Test_CanJumpFail(t *testing.T) {
	source := []int{3, 2, 1, 0, 4}
	target := canJump(source)
	expect := false

	if expect != target {
		t.Error("Translate Jump Game CanFail Fail ", target)
	}
	t.Log("Translate Test_CanJumpFail more Success")
}
