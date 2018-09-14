package weekly_contest_100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTreeNode(list []int) *TreeNode {
	if len(list) < 1 {
		return nil
	}

	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	current := root
	first := &TreeNode{
		Val:   list[0],
		Left:  nil,
		Right: nil,
	}
	current.Right = first

	nodeList := make([]*TreeNode, len(list))
	nodeList[0] = first

	for i := 1; i < len(list); i++ {
		val := list[i]
		if val < 0 {
			nodeList[i] = nil
			continue
		}

		parent := (i - 1) / 2
		point := &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}

		println(i, val, parent, list[parent])
		parentNode := nodeList[parent]
		if 1%2 == 1 {
			parentNode.Left = point
		} else {
			parentNode.Right = point
		}
		nodeList[i] = point
	}

	i*2 - 1 = (11 + 1) / 2
	i * 2 = 12 / 2

	return root.Right
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	return nil
}
