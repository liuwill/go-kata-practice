package weekly_contest_102

import (
	"strings"
	"testing"
)

var (
	sourcesParity = [][]int{
		{3, 1, 2, 4},
	}
	expectsParity = [][]int{
		{2, 4, 3, 1},
	}
)

func compareEvenFirst(expect []int, target []int) bool {
	if len(expect) != len(target) {
		return false
	}

	expectLetters := make([]string, len(expect))
	targetLetters := make([]string, len(expect))
	for i, v := range expect {
		expectLetters[i] = string(v)
		targetLetters[i] = string(target[i])
	}

	expectLine := strings.Join(expectLetters, ",")
	targetLine := strings.Join(targetLetters, ",")

	mark := true
	for i, v := range target {
		if !strings.Contains(expectLine, targetLetters[i]) {
			return false
		}

		if !strings.Contains(targetLine, expectLetters[i]) {
			return false
		}

		if mark && v%2 > 0 {
			return false
		} else if !mark && v%2 == 0 {
			return false
		} else if v%2 > 0 {
			mark = false
		}
	}

	return true
}

func Test_SortArrayByParity(t *testing.T) {
	for i, source := range sourcesParity {
		target := sortArrayByParity(source)
		expect := expectsParity[i]

		if !compareEvenFirst(expect, target) {
			t.Error("Run Test_SortArrayByParityGroup Fail", source, target)
		}
	}
	t.Log("Run Test_SortArrayByParityGroup Success")
}
