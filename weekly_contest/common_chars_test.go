package weekly_contest

import (
	"sort"
	"testing"
)

func compareListWithoutOrder(source []string, target []string) bool {
	if len(source) != len(target) {
		return false
	}
	sort.Strings(source)
	sort.Strings(target)

	for i, letter := range source {
		if target[i] != letter {
			return false
		}
	}
	return true
}

func Test_CommonChars(t *testing.T) {
	sourceCase := [][]string{
		{"bella", "label", "roller"},
		{"cool", "lock", "cook"},
		{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"},
	}
	expectList := [][]string{
		{"e", "l", "l"},
		{"c", "o"},
		{},
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := commonChars(source)
		if !compareListWithoutOrder(target, expect) {
			t.Error("Run Test_CommonChars Fail", expect, target)
		}
	}

	t.Log("Run Test_CommonChars Success")
}
