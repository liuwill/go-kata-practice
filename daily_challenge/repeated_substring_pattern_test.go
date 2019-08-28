package daily_challenge

import "testing"

func Test_RepeatedSubstringPattern(t *testing.T) {
	sourceCase := []string{
		"abab", "aba", "abcabcabcabc", "",
	}
	expectCase := []bool{
		true, false, true, false,
	}

	for i, source := range sourceCase {
		expect := expectCase[i]

		target := repeatedSubstringPattern(source)
		if expect != target {
			t.Error("Run Test_RepeatedSubstringPattern Fail", target)
		}

		targetFast := repeatedSubstringPatternFast(source)
		if expect != targetFast {
			t.Error("Run Test_RepeatedSubstringPatternFast Fail", targetFast)
		}
	}

	t.Log("Run Test_RepeatedSubstringPattern Success")
}
