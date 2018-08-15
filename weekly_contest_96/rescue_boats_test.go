package weekly_contest_96

import (
	"testing"
)

func Test_NumRescueBoats(t *testing.T) {
	source := []int{3, 2, 2, 1}
	limit := 3
	target := numRescueBoats(source, limit)
	expect := 3

	if expect != target {
		t.Error("Translate Test_NumRescueBoats Fail", target)
	}
	t.Log("Translate Test_NumRescueBoats Success")
}

func Test_NumRescueBoatsLess(t *testing.T) {
	source := []int{1, 2}
	limit := 3
	target := numRescueBoats(source, limit)
	expect := 1

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsLess Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsLess Success")
}

func Test_NumRescueBoatsMore(t *testing.T) {
	source := []int{3, 5, 3, 4}
	limit := 5
	target := numRescueBoats(source, limit)
	expect := 4

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsMore Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsMore Success")
}

func Test_NumRescueBoatsTimeout(t *testing.T) {
	source := []int{5, 1, 4, 2}
	limit := 6
	target := numRescueBoats(source, limit)
	expect := 2

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsTimeout Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsTimeout Success")
}

func Test_NumRescueBoatsRuntime(t *testing.T) {
	source := []int{6, 3, 5, 6, 2, 3}
	limit := 6
	target := numRescueBoats(source, limit)
	expect := 5

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsRuntime Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsRuntime Success")
}
