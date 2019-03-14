package weekly_contest

import (
	"testing"
)

func buildBinaryTree(list []int) *TreeNode {
	if len(list) < 1 {
		return nil
	}
	nodeList := make([]*TreeNode, len(list))
	top := 1
	tail := 1

	root := &TreeNode{
		Val:   list[0],
		Left:  nil,
		Right: nil,
	}
	nodeList[0] = root
	parent := root
	mode := false
	for i := 1; i < len(list); i++ {
		val := list[i]
		newNode := &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
		if val == -1 {
			newNode = nil
		}

		nodeList[tail] = newNode
		tail++

		if mode == false {
			parent.Left = newNode
			mode = true
		} else {
			parent.Right = newNode
			for ; top < tail; top++ {
				if nodeList[top] != nil {
					parent = nodeList[top]
					top++
					break
				}
			}
			mode = false
		}
	}

	return root
}

func Test_IsCousins(t *testing.T) {
	rootCase := [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, -1, 4, -1, 5},
		{1, 2, 3, -1, 4},
		{1, 2, 3, -1, 4},
	}
	sourceCase := [][]int{
		{4, 3},
		{5, 4},
		{3, 4},
		{2, 3},
	}
	expectList := []bool{
		false,
		true,
		false,
		false,
	}

	for i, source := range sourceCase {
		root := rootCase[i]
		expect := expectList[i]
		treeRoot := buildBinaryTree(root)

		target := isCousins(treeRoot, source[0], source[1])
		if target != expect {
			t.Error("Run Test_IsCousins Fail", expect, target, root)
		}
	}

	t.Log("Run Test_IsCousins Success")
}
