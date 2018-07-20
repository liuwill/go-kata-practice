package sort_colors

import (
	"testing"
)

func isArrayMatch(target []int, expect []int) bool {
	for i := 0; i < len(target); i++ {
		if target[i] != expect[i] {
			return false
		}
	}
	return true
}

func Test_Sort(t *testing.T) {
	source := []int{2, 0, 2, 1, 1, 0}
	sortColors(source)
	expect := []int{0, 0, 1, 1, 2, 2}

	if !isArrayMatch(source, expect) {
		t.Error("Translate sortColors Fail", source)
	}
	t.Log("Translate Test_Sort Success")
}

func Test_SortError(t *testing.T) {
	source := []int{1, 2, 0}
	sortColors(source)
	expect := []int{0, 1, 2}

	if !isArrayMatch(source, expect) {
		t.Error("Translate sortColors error Fail", source)
	}
	t.Log("Translate Test_SortError Success")
}
