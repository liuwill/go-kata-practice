package weekly_contest

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * daily-challenge-993
 * PUZZLE: Cousins in Binary Tree
 */

func isCousins(root *TreeNode, x int, y int) bool {
	treeStack := []*TreeNode{
		root,
	}

	for len(treeStack) > 0 {
		top := 0
		currentLevel := []*TreeNode{}
		mark := []int{-1, -1}
		for ; top < len(treeStack); top++ {
			current := treeStack[top]
			if current.Left != nil {
				currentLevel = append(currentLevel, current.Left)
			}

			if current.Right != nil {
				currentLevel = append(currentLevel, current.Right)
			}

			if current.Val == x {
				mark[0] = top
			}
			if current.Val == y {
				mark[1] = top
			}
		}

		if mark[0] != mark[1] && mark[0] >= 0 && mark[1] >= 0 {
			return true
		}
		treeStack = currentLevel
	}
	return false
}
