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
		t.Error("Translate Test_Multiply x Fail ", bigTarget)
	}
	t.Log("Translate Test_Multiply Success")
}

func Test_MultiplyBig(t *testing.T) {
	target := multiply("63456346", "6453745")
	expect := "409531075715770"

	if expect != target {
		t.Error("Translate Test_MultiplyBig Fail ", target)
	}

	t.Log("Translate Test_MultiplyBig Success")
}

func Test_MultiplyLarge(t *testing.T) {
	target := multiply("409531075715770", "34534252365264")
	expect := "14142849520186440394285013280"

	if expect != target {
		t.Error("Translate Test_MultiplyLarge Fail ", target)
	}

	t.Log("Translate Test_MultiplyLarge Success")
}

func Test_MultiplyZero(t *testing.T) {
	target := multiply("409531075715770", "0")
	expect := "0"

	if expect != target {
		t.Error("Translate Test_MultiplyZero Fail ", target)
	}

	t.Log("Translate Test_MultiplyZero Success")
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
	target := addStrings("1242", "63")
	expect := "1305"

	if expect != target {
		t.Error("Translate Test_Plus Fail ", target)
	}

	xTarget := addStrings("99999999", "11")
	xExpect := "100000010"

	if xExpect != xTarget {
		t.Error("Translate Test_Plus Fail ", xTarget)
	}

	t.Log("Translate Test_Plus Success")
}

func Test_PlusOne(t *testing.T) {
	target := addStrings("1242", "1")
	expect := "1243"

	if expect != target {
		t.Error("Translate Test_PlusOne Fail ", target)
	}

	t.Log("Translate Test_PlusOne Success")
}

func Test_PlusZero(t *testing.T) {
	target := addStrings("1242", "0")
	expect := "1242"

	if expect != target {
		t.Error("Translate Test_PlusZero Fail ", target)
	}

	t.Log("Translate Test_PlusZero Success")
}
