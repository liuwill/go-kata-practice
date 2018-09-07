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

func Test_LongestValidParenthesesEmpty(t *testing.T) {
	source := ""
	target := longestValidParentheses(source)
	expect := 0

	if expect != target {
		t.Error("Translate Test_LongestValidParenthesesEmpty Fail", target)
	}
	t.Log("Translate Test_LongestValidParenthesesEmpty Success")
}

func Test_LongestValidParenthesesError(t *testing.T) {
	source := "()(()"
	target := longestValidParentheses(source)
	expect := 2

	if expect != target {
		t.Error("Translate Test_LongestValidParenthesesError Fail", target)
	}
	t.Log("Translate Test_LongestValidParenthesesError Success")
}

func Test_LongestValidParenthesesLong(t *testing.T) {
	source := "()((())"
	target := longestValidParentheses(source)
	expect := 4

	if expect != target {
		t.Error("Translate Test_LongestValidParenthesesLong Fail", target)
	}
	t.Log("Translate Test_LongestValidParenthesesLong Success")
}

func Test_LongestValidParenthesesFour(t *testing.T) {
	source := ")()())"
	target := longestValidParentheses(source)
	expect := 4

	if expect != target {
		t.Error("Translate Test_LongestValidParenthesesFour Fail", target)
	}
	t.Log("Translate Test_LongestValidParenthesesFour Success")
}

func Test_LongestValidParenthesesCase(t *testing.T) {
	source := "()(())))()(()("
	target := longestValidParentheses(source)
	expect := 6

	if expect != target {
		t.Error("Translate Test_LongestValidParenthesesCase Fail", target)
	}
	t.Log("Translate Test_LongestValidParenthesesCase Success")
}

func Test_LongestValidParenthesesMore(t *testing.T) {
	source := "()(())))()(()())"
	target := longestValidParentheses(source)
	expect := 8

	if expect != target {
		t.Error("Translate Test_LongestValidParenthesesMore Fail", target)
	}
	t.Log("Translate Test_LongestValidParenthesesMore Success")
}
