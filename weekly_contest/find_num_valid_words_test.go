package weekly_contest

import "testing"

func Test_FindNumOfValidWords(t *testing.T) {
	wordCase := [][]string{
		{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
	}
	puzzleList := [][]string{
		{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"},
	}
	expectList := [][]int{
		{1, 1, 3, 2, 4, 0},
	}

	for i, word := range wordCase {
		expect := expectList[i]
		puzzle := puzzleList[i]

		target := findNumOfValidWords(word, puzzle)
		if !compareList(target, expect) {
			t.Error("Run Test_FindNumOfValidWords Fail", expect, target)
		}
	}

	t.Log("Run Test_FindNumOfValidWords Success")
}
