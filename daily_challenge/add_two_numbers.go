package daily_challenge

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	target := &ListNode{
		Val:  0,
		Next: nil,
	}

	first := l1
	second := l2

	front := 0
	result := target
	for first != nil || second != nil {
		num1 := 0
		num2 := 0
		if first != nil {
			num1 = first.Val
			first = first.Next
		}

		if second != nil {
			num2 = second.Val
			second = second.Next
		}

		fmt.Printf("%v %v %v", l1, second, first)

		sum := num1 + num2 + front
		current := sum % 10
		front = sum / 10

		currentNode := &ListNode{
			Val:  current,
			Next: nil,
		}
		target.Next = currentNode
		target = currentNode
	}

	return result
}
