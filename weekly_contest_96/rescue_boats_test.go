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
