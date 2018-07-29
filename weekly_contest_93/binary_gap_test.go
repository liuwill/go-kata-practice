package weekly_contest_93

import (
	"testing"
)

func Test_BinaryGap(t *testing.T) {
	source := 22
	target := binaryGap(source)
	expect := 2

	if expect != target {
		t.Error("Translate binaryGap Fail", target)
	}
	t.Log("Translate Test_BinaryGap Success")
}

func Test_BinaryGapZero(t *testing.T) {
	source := 8
	target := binaryGap(source)
	expect := 0

	if expect != target {
		t.Error("Translate binaryGap zero Fail", target)
	}
	t.Log("Translate Test_BinaryGapZero Success")
}

func Test_BinaryGapFive(t *testing.T) {
	source := 5
	target := binaryGap(source)
	expect := 2

	if expect != target {
		t.Error("Translate binaryGap five Fail", target)
	}
	t.Log("Translate Test_BinaryGapFive Success")
}

func Test_BinaryGapSix(t *testing.T) {
	source := 6
	target := binaryGap(source)
	expect := 1

	if expect != target {
		t.Error("Translate binaryGap six Fail", target)
	}
	t.Log("Translate Test_BinaryGapSix Success")
}
