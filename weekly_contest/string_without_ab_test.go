package weekly_content

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
			size = 0
			current = letter
		}

		if size >= 3 {
			return false
		}
	}
	return true
}

func Test_StrWithout3a3b(t *testing.T) {
	source := [][]int{
		{1, 2},
		{4, 1},
	}

	for _, pair := range source {
		target := strWithout3a3b(pair[0], pair[1])
		if !checkStrWithout3a3b(target, pair[0], pair[1]) {
			t.Error("Translate Test_StrWithout3a3b Fail", target)
		}
	}

	t.Log("Translate Test_StrWithout3a3b Success")
}