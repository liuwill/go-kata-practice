package weekly_contest_100

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

func Test_GenerateTreeNode(t *testing.T) {
	// source := []int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9}
	sourceTree := generateTreeNode([]int{})

	if sourceTree != nil {
		t.Error("Translate Test_GenerateTreeNode Fail ")
	}
	t.Log("Translate Test_GenerateTreeNode Success")
}

func Test_RecoverTreeNode(t *testing.T) {
	targetList := recoverTreeNode(nil)

	if len(targetList) != 0 {
		t.Error("Translate Test_RecoverTreeNode Fail ")
	}

	source := []int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9}
	sourceTree := generateTreeNode(source)
	target := recoverTreeNode(sourceTree)

	if !compareList(source, target) {
		t.Error("Translate Test_RecoverTreeNode Fail ", target)
	}

	t.Log("Translate Test_RecoverTreeNode Success")
}

func Test_IncreasingBST(t *testing.T) {
	source := []int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9}
	sourceTree := generateTreeNode(source)
	rawTarget := increasingBST(sourceTree)
	target := recoverTreeNode(rawTarget)
	expect := []int{1, -1, 2, -1, 3, -1, 4, -1, 5, -1, 6, -1, 7, -1, 8, -1, 9}

	if !compareList(expect, target) {
		t.Error("Translate Test_IncreasingBST Fail ", target)
	}
	t.Log("Translate Test_IncreasingBST Success")
}
