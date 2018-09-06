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

func transformNumberToListNode(number int) *ListNode {
	if number == 0 {
		return &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	current := &ListNode{
		Val:  -1,
		Next: nil,
	}

	result := current
	for number > 0 {
		bit := number % 10
		number = number / 10

		if current.Val == -1 {
			current.Val = bit
		} else {
			current.Next = &ListNode{
				Val:  bit,
				Next: nil,
			}
			current = current.Next
		}
	}

	return result
}

func printListNode(list *ListNode) []int {
	result := []int{}

	current := list
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

func Test_TransformNumberToListNode(t *testing.T) {
	list := transformNumberToListNode(342)
	target := printListNode(list)
	expect := []int{2, 4, 3}

	if !compareList(expect, target) {
		t.Error("Translate Test_TransformNumberToListNode Fail", target)
	}
}

func Test_AddTwoNumbers(t *testing.T) {
	source1 := transformNumberToListNode(0)
	source2 := transformNumberToListNode(0)
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
