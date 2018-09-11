package daily_challenge

import "testing"

func Test_ScoreOfParentheses(t *testing.T) {
	source := "()"
	target := scoreOfParentheses(source)
	expect := 1

	if expect != target {
		t.Error("Translate Test_ScoreOfParentheses Fail", target)
	}
	t.Log("Translate Test_ScoreOfParentheses Success")
}

func Test_ScoreOfParenthesesTwo(t *testing.T) {
	source := "()()"
	target := scoreOfParentheses(source)
	expect := 2

	if expect != target {
		t.Error("Translate Test_ScoreOfParenthesesTwo Fail", target)
	}
	t.Log("Translate Test_ScoreOfParenthesesTwo Success")
}
