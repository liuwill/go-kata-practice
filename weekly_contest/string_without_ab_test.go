package weekly_contest

import (
	"testing"
)

func checkStrWithout3a3b(str string, A int, B int) bool {
	if len(str) != A+B {
		return false
	}

	current := 'x'
	size := 0
	for _, letter := range str {
		if current == letter {
			size++
		} else {
			size = 1
			current = letter
		}

		if size >= 3 {
			return false
		}
	}

	return size < 3
}

func Test_StrWithout3a3b(t *testing.T) {
	source := [][]int{
		{1, 2},
		{4, 1},
		{0, 0},
		{2, 5},
		{2, 6},
		{8, 3},
	}

	for _, pair := range source {
		target := strWithout3a3b(pair[0], pair[1])
		if !checkStrWithout3a3b(target, pair[0], pair[1]) {
			println("A:", pair[0], "B:", pair[1], "=>", target)
			t.Error("Translate Test_StrWithout3a3b Fail", target)
		}
	}

	t.Log("Translate Test_StrWithout3a3b Success")
}
