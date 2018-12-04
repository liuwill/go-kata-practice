package daily_challenge

import (
	"testing"
)

func Test_MinIncrementForUniqueEasy(t *testing.T) {
	source := []int{1, 2, 2}
	expect := 1

	target := minIncrementForUnique(source)

	if target != expect {
		t.Error("Translate Test_MinIncrementForUniqueEasy Fail", target, expect)
	}
	t.Log("Translate Test_MinIncrementForUniqueEasy Success")
}

func Test_MinIncrementForUnique(t *testing.T) {
	source := []int{3, 2, 1, 2, 1, 7}
	expect := 6

	target := minIncrementForUnique(source)

	if target != expect {
		t.Error("Translate Test_MinIncrementForUnique Fail", target, expect)
	}
	t.Log("Translate Test_MinIncrementForUnique Success")
}
