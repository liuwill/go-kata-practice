package daily_challenge

import (
	"testing"
)

func compareListNode(expect *ListNode, target *ListNode) bool {
	first := expect
	second := target
	for first != nil && second != nil {
		if first == nil {
			return false
		} else if second == nil {
			return false
		}

		if first.Val != second.Val {
			return false
		}

		first = first.Next
		second = second.Next
	}

	return true
}

func Test_AddTwoNumbers(t *testing.T) {
	source1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	source2 := &ListNode{
		Val:  0,
		Next: nil,
	}
	target := addTwoNumbers(source1, source2)
	expect := &ListNode{
		Val:  0,
		Next: nil,
	}

	if !compareListNode(expect, target) {
		t.Error("Translate Test_AddTwoNumbers Fail", target)
	}
	t.Log("Translate Test_AddTwoNumbers Success")
}
