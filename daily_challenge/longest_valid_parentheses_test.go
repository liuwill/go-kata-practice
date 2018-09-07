package daily_challenge

import "testing"

func Test_LongestValidParentheses(t *testing.T) {
	source := "(()"
	target := longestValidParentheses(source)
	expect := 2

	if expect != target {
		t.Error("Translate Test_LongestValidParentheses Fail", target)
	}
	t.Log("Translate Test_LongestValidParentheses Success")
}
