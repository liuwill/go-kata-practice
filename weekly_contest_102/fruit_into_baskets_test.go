package weekly_contest_102

import "testing"

func Test_TotalFruit(t *testing.T) {
	source := []int{1, 2, 2, 3}
	target := totalFruit(source)
	expect := 3

	if expect != target {
		t.Error("Run Test_TotalFruit Fail", target)
	}
	t.Log("Run Test_TotalFruit Success")
}
