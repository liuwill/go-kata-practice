package multiply_strings

import (
	"testing"
)

func Test_Multiply(t *testing.T) {
	target := multiply("2", "3")
	expect := "6"

	if expect != target {
		t.Error("Translate Test_Multiply Fail ", target)
	}

	bigTarget := multiply("25", "25")
	bigExpect := "625"
	if bigExpect != bigTarget {
		t.Error("Translate Test_Multiply Fail ", bigTarget)
	}
	t.Log("Translate Test_Multiply Success")
}
