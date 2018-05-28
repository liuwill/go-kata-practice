package swap_pair_node

import (
	"testing"
)

func Test_Swap(t *testing.T) {
	source := []int{1, 2, 3, 4}
	sourceNodes := buildPairList(source)
	targetNodes := swapPairs(sourceNodes)
	target := parseListNode(targetNodes)
	expect := []int{2, 1, 4, 3}

	if !compareList(expect, target) {
		t.Error("Translate swap node Fail ", target, parseListNode(sourceNodes), []int{1, 2, 3, 4})
	}
	t.Log("Translate Test_Swap Success")
}
