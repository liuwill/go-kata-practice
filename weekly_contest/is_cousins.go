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

	source := []int{x, y}
	for len(treeStack) > 0 {
		top := 0
		currentLevel := []*TreeNode{}
		mark := []int{-1, -1}
		for ; top < len(treeStack); top++ {
			current := treeStack[top]
			if current.Left != nil {
				currentLevel = append(currentLevel, current.Left)
				for i, v := range source {
					if current.Left.Val == v {
						mark[i] = top
					}
				}
			}

			if current.Right != nil {
				currentLevel = append(currentLevel, current.Right)
				for i, v := range source {
					if current.Right.Val == v {
						mark[i] = top
					}
				}
			}
		}

		if mark[0] != mark[1] && mark[0] >= 0 && mark[1] >= 0 {
			return true
		}
		treeStack = currentLevel
	}
	return false
}
