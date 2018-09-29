package weekly_contest_100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTreeNode(tree *TreeNode) []int {
	if tree == nil {
		return []int{}
	}

	nodeList := []*TreeNode{
		tree,
	}
	result := []int{
		tree.Val,
	}
	cursor := 1

	pos := 0
	length := 1
	for pos < cursor {
		current := nodeList[pos]

		if current.Left != nil {
			result = append(result, current.Left.Val)
			nodeList = append(nodeList, current.Left)
			cursor++

			length = len(result)
		} else {
			result = append(result, -1)
		}

		if current.Right != nil {
			result = append(result, current.Right.Val)
			nodeList = append(nodeList, current.Right)
			cursor++

			length = len(result)
		} else {
			result = append(result, -1)
		}

		pos++
	}

	return result[:length]
}

func generateTreeNode(list []int) *TreeNode {
	if len(list) < 1 || list[0] < 0 {
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
	resultTree := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	nodeList := []*TreeNode{}
	cursor := 0

	current := root
	target := resultTree

	pushStack := func(stack []*TreeNode, node *TreeNode, cursor int) []*TreeNode {
		if cursor < len(stack) {
			stack[cursor] = node
		} else {
			stack = append(stack, node)
		}
		return stack
	}

	for current != nil || cursor > 0 {
		if current.Left != nil {
			nodeList = pushStack(nodeList, current, cursor)

			next := current
			current = next.Left
			next.Left = nil
			cursor++
			continue
		}

		target.Right = &TreeNode{
			Val:   current.Val,
			Left:  nil,
			Right: nil,
		}
		target = target.Right

		if current.Right != nil {
			nodeList = pushStack(nodeList, current.Right, cursor)

			next := current
			current = next.Right
			next.Right = nil

			cursor++
		}

		if cursor > 0 {
			current = nodeList[cursor-1]
		} else {
			current = nil
		}
		cursor--
	}

	return resultTree.Right
}
