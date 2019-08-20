package weekly_contest

import "testing"

func Test_LastSubstring(t *testing.T) {
	mapCase := []string{
		"abab",
		"leetcode",
	}
	expectList := []string{
		"bab",
		"tcode",
	}

	for i, source := range mapCase {
		expect := expectList[i]

		target := lastSubstring(source)
		if target != expect {
			t.Error("Run Test_LastSubstring Fail", expect, target)
		}
	}

	t.Log("Run Test_LastSubstring Success")
}
