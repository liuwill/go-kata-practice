package daily_challenge

import "testing"

func Test_IsValid(t *testing.T) {
	source := "()"
	target := isValid(source)
	expect := true

	if expect != target {
		t.Error("Translate Test_IsValid Fail", target)
	}
	t.Log("Translate Test_IsValid Success")
}

func Test_IsValidFull(t *testing.T) {
	source := "()[]{}"
	target := isValid(source)
	expect := true

	if expect != target {
		t.Error("Translate Test_IsValidFull Fail", target)
	}
	t.Log("Translate Test_IsValidFull Success")
}

func Test_IsValidFalse(t *testing.T) {
	source := "(]"
	target := isValid(source)
	expect := false

	if expect != target {
		t.Error("Translate Test_IsValidFalse Fail", target)
	}
	t.Log("Translate Test_IsValidFalse Success")
}

func Test_IsValidFalseMore(t *testing.T) {
	source := "([)]"
	target := isValid(source)
	expect := false

	if expect != target {
		t.Error("Translate Test_IsValidFalseMore Fail", target)
	}
	t.Log("Translate Test_IsValidFalseMore Success")
}

func Test_IsValidTrue(t *testing.T) {
	source := "{[]}"
	target := isValid(source)
	expect := true

	if expect != target {
		t.Error("Translate Test_IsValidTrue Fail", target)
	}
	t.Log("Translate Test_IsValidTrue Success")
}

func Test_IsValidEmpty(t *testing.T) {
	source := ""
	target := isValid(source)
	expect := true

	if expect != target {
		t.Error("Translate Test_IsValidEmpty Fail", target)
	}
	t.Log("Translate Test_IsValidEmpty Success")
}
