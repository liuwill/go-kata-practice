package weekly_contest

import (
	"testing"
)

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

func Test_FindOccurrences(t *testing.T) {
	sourceCase := [][]string{
		[]string{
			"alice is a good girl she is a good student",
			"a",
			"good",
		},
		[]string{
			"we will we will rock you",
			"we",
			"will",
		},
	}
	expectList := [][]string{
		[]string{"girl", "student"},
		[]string{"we", "rock"},
	}

	for i, source := range sourceCase {
		expect := expectList[i]

		target := findOcurrences(source[0], source[1], source[2])
		if !compareStringList(target, expect) {
			t.Error("Run Test_FindOcurrences Fail", expect, target)
		}
	}

	t.Log("Run Test_FindOcurrences Success")
}
