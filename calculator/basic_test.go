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

func Test_CanCalculateBasicLong(t *testing.T) {
	source := "3 + 5 * 2 - 4 "

	target := calculate(source)
	expect := 9

	if expect != target {
		t.Error("Translate Calculate Long CanSuccess Fail ", target)
	}
	t.Log("Translate Test_CanCalculateBasicLong Success")
}

func Test_CanCalculateBasicWrong(t *testing.T) {
	source := "1-1+1"

	target := calculate(source)
	expect := 1

	if expect != target {
		t.Error("Translate Calculate Wrong CanSuccess Fail ", target)
	}
	t.Log("Translate Test_CanCalculateBasicWrong Success")
}
