package main

import (
	"fmt"
	. "github.com/fyibmsd/interview/common"
)

/*
 * Q: 从尾到头打印单链表
 *
 * 思路：利用栈
 */
func printListReversingly(head *ListNode) {
	stack := NewStack()
	curr := head

	for curr != nil {
		stack.Push(curr.Data)

		curr = curr.Next
	}

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}

func main() {
	list := NewLinkedList([]int{1, 3, 5, 7, 9})

	printListReversingly(list)
}
