package calculator

import (
	"testing"
)

func Test_CanCalculateBasicSuccess(t *testing.T) {
	source := "3+2*2"

	target := calculate(source)
	expect := 7

	if expect != target {
		t.Error("Translate Calculate Basic CanSuccess Fail ", target)
	}
	t.Log("Translate Test_CanCalculateBasicSuccess Success")
}
