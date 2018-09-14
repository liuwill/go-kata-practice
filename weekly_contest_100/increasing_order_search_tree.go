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

	first := &TreeNode{
		Val:   list[0],
		Left:  nil,
		Right: nil,
	}

	nodeList := make([]*TreeNode, len(list))
	nodeList[0] = first

	cursor := 1
	parent := 0
	for i := 1; i < len(list) && parent < cursor; i += 2 {
		left := list[i]
		right := list[i+1]

		currentNode := nodeList[parent]
		if left >= 0 {
			leftNode := &TreeNode{
				Val:   left,
				Left:  nil,
				Right: nil,
			}
			currentNode.Left = leftNode
			nodeList[cursor] = leftNode
			cursor++
		} else {
			currentNode.Left = nil
		}

		if right >= 0 {
			rightNode := &TreeNode{
				Val:   right,
				Left:  nil,
				Right: nil,
			}
			currentNode.Right = rightNode
			nodeList[cursor] = rightNode
			cursor++
		} else {
			currentNode.Right = nil
		}

		parent++
	}

	// fmt.Printf("-=> %v\n", nodeList)
	// for i, v := range nodeList {
	// 	fmt.Printf("%v - %v\n", i, v)
	// }

	return first
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
