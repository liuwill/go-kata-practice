package count_say

import (
	"testing"
)

func Test_CountSay(t *testing.T) {
	source := 1
	target := countAndSay(source)
	expect := "1"

	if expect != target {
		t.Error("Translate CountAndSay Fail ", target)
	}
	t.Log("Translate Test_CountSay Success")
}

func Test_CountSayFour(t *testing.T) {
	source := 4
	target := countAndSay(source)
	expect := "1211"

	if expect != target {
		t.Error("Translate CountAndSay four Fail ", target)
	}
	t.Log("Translate Test_CountSay four Success")
}
