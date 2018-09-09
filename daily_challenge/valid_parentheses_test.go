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
