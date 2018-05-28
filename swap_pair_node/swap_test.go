package swap_pair_node

import (
	"testing"
)

func Test_Length(t *testing.T) {
	source := []int{1, 2, 3, 4}
	sourceNodes := buildPairList(source)
	target := lenPairList(sourceNodes)
	expect := 4

	if expect != target {
		t.Error("Translate swap lenPairList Fail ", target)
	}
	t.Log("Translate test_Length Success")
}

func Test_Compare(t *testing.T) {
	source := []int{1, 2, 3, 4}
	target := []int{1, 2, 4}
	expect := []int{1, 2, 3}

	if compareList(expect, target) {
		t.Error("Translate swap compareList Fail ", target)
	}

	if compareList(source, target) {
		t.Error("Translate swap compareList Fail ", target)
	}
	t.Log("Translate test_Compare Success")
}

func Test_Swap(t *testing.T) {
	source := []int{1, 2, 3, 4}
	sourceNodes := buildPairList(source)
	targetNodes := swapPairs(sourceNodes)
	target := parseListNode(targetNodes)
	expect := []int{2, 1, 4, 3}

	if !compareList(expect, target) {
		t.Error("Translate swap node Fail ", target)
	}
	t.Log("Translate Test_Swap Success")
}

func Test_SwapError(t *testing.T) {
	source := []int{2, 5, 3, 4, 6, 2, 2}
	sourceNodes := buildPairList(source)
	targetNodes := swapPairs(sourceNodes)
	target := parseListNode(targetNodes)
	expect := []int{5, 2, 4, 3, 2, 6, 2}

	if !compareList(expect, target) {
		t.Error("Translate swap node error Fail ", target)
	}
	t.Log("Translate Test_SwapError Success")
}
