package weekly_contest_98

import "testing"

func Test_SumSubseqWidths(t *testing.T) {
	source := []int{2, 1, 3}
	target := sumSubseqWidths(source)
	expect := 6

	if expect != target {
		t.Error("Translate Test_SumSubseqWidths Fail", expect, target)
	}
	t.Log("Translate Test_SumSubseqWidths Success")
}
