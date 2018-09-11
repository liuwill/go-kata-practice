package weekly_contest_100

import "testing"

func Test_IsMonotonic(t *testing.T) {
	source := []int{1, 2, 2, 3}
	target := isMonotonic(source)
	expect := true

	if expect != target {
		t.Error("Translate Test_IsMonotonic Fail", target)
	}
	t.Log("Translate Test_IsMonotonic Success")
}
