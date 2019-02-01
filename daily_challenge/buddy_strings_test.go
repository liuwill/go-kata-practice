package daily_challenge

import "testing"

func Test_BuddyStrings(t *testing.T) {
	source := [][]string{
		{"ab", "ba"},
		{"ab", "ab"},
		{"aa", "aa"},
		{"aaaaaaabc", "aaaaaaacb"},
		{"", "aa"},
	}
	expect := []bool{
		true,
		false,
		true,
		true,
		false,
	}

	for i, val := range source {
		target := buddyStrings(val[0], val[1])
		if expect[i] != target {
			t.Error("Translate Test_BuddyStrings Fail", val, target)
		}
	}

	t.Log("Translate Test_BuddyStrings Success")
}
