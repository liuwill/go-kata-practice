package remove_duplicates

import (
	"testing"
)

func Test_Remove(t *testing.T) {
	source := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	println("++++++", source)
	target := removeDuplicates(source[:])
	expect := 5

	if expect != target || len(source) != expect {
		t.Error("Translate remove duplicates Fail ", target, len(source), source)
	}
	t.Log("Translate Test_Remove Success")
}
