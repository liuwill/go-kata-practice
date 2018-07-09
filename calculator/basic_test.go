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

func Test_CanCalculateBasicBlank(t *testing.T) {
	source := "3+5 / 2"

	target := calculate(source)
	expect := 5

	if expect != target {
		t.Error("Translate Calculate Blank CanSuccess Fail ", target)
	}
	t.Log("Translate Test_CanCalculateBasicBlank Success")
}
