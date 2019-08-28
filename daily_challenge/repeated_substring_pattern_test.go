package daily_challenge

import "testing"

func Test_RepeatedSubstringPattern(t *testing.T) {
	sourceCase := []string{
		"abab", "aba", "abcabcabcabc",
	}
	expectCase := []bool{
		true, false, true,
	}

	for i, source := range sourceCase {
		expect := expectCase[i]

		target := repeatedSubstringPattern(source)
		if expect != target {
			t.Error("Run Test_RepeatedSubstringPattern Fail", target)
		}
	}

	t.Log("Run Test_RepeatedSubstringPattern Success")
}
