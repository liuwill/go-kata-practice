package remove_duplicates

import (
	"testing"
)

func compareList(source []int, target []int) bool {
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

func Test_Remove(t *testing.T) {
	source := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	target := removeDuplicates(source[:])
	truth := []int{0, 1, 2, 3, 4}
	expect := 5

	if expect != target || !compareList(truth, source[:target]) {
		t.Error("Translate remove duplicates Fail ", target, source[:target])
	}
	t.Log("Translate Test_Remove Success")
}

func Test_RemoveLess(t *testing.T) {
	source := []int{4}
	target := removeDuplicates(source[:])
	expect := 1

	if expect != target {
		t.Error("Translate remove duplicates Fail ", target)
	}
	t.Log("Translate Test_Remove Success")
}
