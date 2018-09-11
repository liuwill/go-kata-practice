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

func Test_ScoreOfParenthesesEmpty(t *testing.T) {
	source := ""
	target := scoreOfParentheses(source)
	expect := 0

	if expect != target {
		t.Error("Translate Test_ScoreOfParenthesesEmpty Fail", target)
	}
	t.Log("Translate Test_ScoreOfParenthesesEmpty Success")
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

func Test_ScoreOfParenthesesLayer(t *testing.T) {
	source := "(())"
	target := scoreOfParentheses(source)
	expect := 2

	if expect != target {
		t.Error("Translate Test_ScoreOfParenthesesLayer Fail", target)
	}
	t.Log("Translate Test_ScoreOfParenthesesLayer Success")
}

func Test_ScoreOfParenthesesDeep(t *testing.T) {
	source := "(()(()))"
	target := scoreOfParentheses(source)
	expect := 6

	if expect != target {
		t.Error("Translate Test_ScoreOfParenthesesDeep Fail", target)
	}
	t.Log("Translate Test_ScoreOfParenthesesDeep Success")
}
