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
		t.Error("Translate CountAndSay Four Fail ", target)
	}
	t.Log("Translate Test_CountSay Four Success")
}

func Test_CountSaySever(t *testing.T) {
	source := 7
	target := countAndSay(source)
	expect := "13112221"

	if expect != target {
		t.Error("Translate CountAndSay Sever Fail ", target)
	}
	t.Log("Translate Test_CountSay Sever Success")
}
