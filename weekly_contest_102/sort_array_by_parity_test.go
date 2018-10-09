package weekly_contest_102

import (
	"strings"
	"testing"
)

var (
	sourcesParity = [][]int{
		{2},
		{1},
		{3, 1, 2, 4},
		{1, 2, 5, 4, 6, 7, 8, 3},
	}
	expectsParity = [][]int{
		{2},
		{1},
		{2, 4, 3, 1},
		{8, 2, 6, 4, 5, 7, 3, 1},
	}
)

func compareEvenFirst(source []int, expect []int, target []int) bool {
	if len(source) != len(target) || len(expect) != len(target) {
		return false
	}

	expectMatch := true
	for i := 0; i < len(target); i++ {
		if target[i] != expect[i] {
			expectMatch = false
			break
		}
	}

	if expectMatch {
		return true
	}

	sourceLetters := make([]string, len(source))
	targetLetters := make([]string, len(source))
	for i, v := range source {
		sourceLetters[i] = string(v)
		targetLetters[i] = string(target[i])
	}

	sourceLine := strings.Join(sourceLetters, ",")
	targetLine := strings.Join(targetLetters, ",")

	mark := true
	for i, v := range target {
		if !strings.Contains(sourceLine, targetLetters[i]) {
			return false
		}

		if !strings.Contains(targetLine, sourceLetters[i]) {
			return false
		}

		if v%2 > 0 {
			mark = false
		} else if !mark && v%2 == 0 {
			return false
		}
	}
	return true
}

func Test_SortArrayByParity(t *testing.T) {
	for i, source := range sourcesParity {
		target := sortArrayByParity(source)
		expect := expectsParity[i]

		if !compareEvenFirst(source, expect, target) {
			t.Error("Run Test_SortArrayByParityGroup Fail", source, target)
		}
	}
	t.Log("Run Test_SortArrayByParityGroup Success")
}
