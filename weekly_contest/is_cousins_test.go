package weekly_contest

import "testing"

func buildBinaryTree(list []int) *TreeNode {
	if len(list) < 1 {
		return nil
	}

	root := &TreeNode{
		Val:   list[0],
		Left:  nil,
		Right: nil,
	}
	parent := root
	mode := false
	for i := 1; i < len(list); i++ {
		val := list[i]
		if mode == false {
			parent.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		} else {
			parent.Right = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			parent = parent.Left
		}

	}

	return root
}

func Test_IsCousins(t *testing.T) {
	rootCase := [][]int{
		{1, 2, 3, 4},
	}
	sourceCase := [][]int{
		{4, 3},
	}
	expectList := []bool{
		false,
	}

	for i, source := range sourceCase {
		root := rootCase[i]
		expect := expectList[i]
		treeRoot := buildBinaryTree(root)

		target := isCousins(treeRoot, source[0], source[1])
		if target != expect {
			t.Error("Run Test_IsCousins Fail", expect, target)
		}
	}

	t.Log("Run Test_IsCousins Success")
}
