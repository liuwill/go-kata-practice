package weekly_contest_100

import (
	"testing"
)

func compareTreeNodeList(expect []int, target *TreeNode) bool {
	return false
}

func Test_IncreasingBST(t *testing.T) {
	// source := []int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9}
	target := increasingBST(nil)
	expect := []int{1, -1, 2, -1, 3, -1, 4, -1, 5, -1, 6, -1, 7, -1, 8, -1, 9}

	if !compareTreeNodeList(expect, target) {
		t.Error("Translate Test_IncreasingBST Fail ", target)
	}
	t.Log("Translate Test_IncreasingBST Success")
}
