package weekly_contest_98

import "testing"

func compareStringList(source []string, target []string) bool {
	if len(source) != len(target) {
		return false
	}

	for index := 0; index < len(source); index++ {
		if source[index] != target[index] {
			return false
		}
	}

	return true
}
func Test_FindAndReplacePattern(t *testing.T) {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	target := findAndReplacePattern(words, pattern)
	expect := []string{"mee", "aqq"}

	if !compareStringList(expect, target) {
		t.Error("Translate Test_FindAndReplacePattern Fail", target)
	}
	t.Log("Translate Test_FindAndReplacePattern Success")
}

func Test_FindAndReplacePatternEmpty(t *testing.T) {
	words := []string{"abc"}
	pattern := "abb"
	target := findAndReplacePattern(words, pattern)
	expect := []string{}

	if !compareStringList(expect, target) {
		t.Error("Translate Test_FindAndReplacePatternEmpty Fail", target)
	}
	t.Log("Translate Test_FindAndReplacePatternEmpty Success")
}
