package weekly_contest

import "testing"

func Test_AlphabetBoardPath(t *testing.T) {
	sourceCase := []string{
		"leet",
		"code",
		"zdz",
	}
	expectList := []string{
		"DDR!UURRR!!DDD!",
		"RR!DDRR!UUL!R!",
		"DDDDD!UUUUURRR!DDDDLLLD!",
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := alphabetBoardPath(source)
		if target != expect {
			t.Error("Run Test_AlphabetBoardPath Fail", expect, target)
		}
	}

	t.Log("Run Test_AlphabetBoardPath Success")
}
