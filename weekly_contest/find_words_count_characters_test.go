package weekly_contest

import "testing"

func Test_CountCharacters(t *testing.T) {
	wordCase := [][]string{
		{"cat", "bt", "hat", "tree"},
		{"hello", "world", "leetcode"},
	}
	charsCase := []string{
		"atach",
		"welldonehoneyr",
	}
	expectList := []int{
		6,
		10,
	}

	for i, words := range wordCase {
		chars := charsCase[i]
		expect := expectList[i]

		target := countCharacters(words, chars)
		if target != expect {
			t.Error("Run Test_CountCharacters Fail", expect, target)
		}
	}

	t.Log("Run Test_CountCharacters Success")
}
