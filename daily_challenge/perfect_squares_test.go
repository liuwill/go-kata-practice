package daily_challenge

import "testing"

func Test_NumSquares(t *testing.T) {
	source := []int{12, 13}
	expect := []int{3, 2}

	for i, val := range source {
		target := numSquares(val)
		if expect[i] != target {
			t.Error("Translate Test_NumSquares Fail", val, target)
		}
	}

	t.Log("Translate Test_NumSquares Success")
}
