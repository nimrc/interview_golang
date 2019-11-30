package common

/*
 * 单链表 Linked list
 */

type ListNode struct {
	Data int
	Next *ListNode
}

func NewLinkedList(values []int) *ListNode {
	head := &ListNode{Data: values[0], Next: nil}
	curr := head

	for _, data := range values[1:] {
		curr.Next = &ListNode{Data: data, Next: nil}
		curr = curr.Next
	}

	return head
}
