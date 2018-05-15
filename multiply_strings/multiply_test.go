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

func Test_Shift(t *testing.T) {
	target := shift("1242", 0)
	expect := "1242"

	if expect != target {
		t.Error("Translate Test_Shift Fail ", target)
	}

	xTarget := shift("432", 4)
	xExpect := "4320000"

	if xExpect != xTarget {
		t.Error("Translate Test_Shift Fail ", xTarget)
	}

	t.Log("Translate Test_Shift Success")
}

func Test_Plus(t *testing.T) {
	target := plus("1242", "63")
	expect := "1305"

	if expect != target {
		t.Error("Translate Test_Plus Fail ", target)
	}

	xTarget := plus("99999999", "11")
	xExpect := "100000010"

	if xExpect != xTarget {
		t.Error("Translate Test_Plus Fail ", xTarget)
	}

	t.Log("Translate Test_Plus Success")
}
