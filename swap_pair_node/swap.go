package swap_pair_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func lenPairList(head *ListNode) int {
	cursor := head
	length := 0
	for cursor != nil {
		length++

		cursor = cursor.Next
	}

	return length
}

func buildPairList(list []int) *ListNode {
	head := &ListNode{
		Val:  list[0],
		Next: nil,
	}

	cursor := head
	for index, value := range list {
		if index == 0 {
			continue
		}

		current := &ListNode{
			Val:  value,
			Next: nil,
		}
		cursor.Next = current
		cursor = current
	}

	return head
}

func parseListNode(head *ListNode) []int {
	cursor := head
	list := []int{}
	for cursor != nil {
		list = append(list, cursor.Val)
		cursor = cursor.Next
	}

	return list
}

func compareList(source []int, target []int) bool {
	if len(source) != len(target) {
		return false
	}

	for index := 0; index < len(source); index++ {
		if source[index] != target[index] {
			return false
		}
	}

	return true
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	cursor := head
	last := head

	index := 0
	for cursor != nil && cursor.Next != nil {
		current := cursor
		next := cursor.Next

		cursor = next.Next
		current.Next = next.Next
		next.Next = current

		if index == 0 {
			last = current
			head = next
		} else {
			last.Next = next
			last = current
		}

		index++
	}

	return head
}
