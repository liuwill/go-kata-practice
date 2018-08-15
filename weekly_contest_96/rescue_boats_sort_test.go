package weekly_contest_96

import (
	"testing"
)

func Test_NumRescueBoatsSort(t *testing.T) {
	source := []int{3, 2, 2, 1}
	limit := 3
	target := numRescueBoatsSort(source, limit)
	expect := 3

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsSort Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsSort Success")
}

func Test_NumRescueBoatsSortLess(t *testing.T) {
	source := []int{1, 2}
	limit := 3
	target := numRescueBoatsSort(source, limit)
	expect := 1

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsSortLess Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsSortLess Success")
}

func Test_NumRescueBoatsSortMore(t *testing.T) {
	source := []int{3, 5, 3, 4}
	limit := 5
	target := numRescueBoatsSort(source, limit)
	expect := 4

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsSortMore Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsSortMore Success")
}

func Test_NumRescueBoatsSortTimeout(t *testing.T) {
	source := []int{5, 1, 4, 2}
	limit := 6
	target := numRescueBoatsSort(source, limit)
	expect := 2

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsSortTimeout Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsSortTimeout Success")
}

func Test_NumRescueBoatsSortRuntime(t *testing.T) {
	source := []int{6, 3, 5, 6, 2, 3}
	limit := 6
	target := numRescueBoatsSort(source, limit)
	expect := 5

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsSortRuntime Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsSortRuntime Success")
}

func Test_NumRescueBoatsSortError(t *testing.T) {
	source := []int{3, 2, 3, 2, 2}
	limit := 6
	target := numRescueBoatsSort(source, limit)
	expect := 3

	if expect != target {
		t.Error("Translate Test_NumRescueBoatsSortError Fail", target)
	}
	t.Log("Translate Test_NumRescueBoatsSortError Success")
}
