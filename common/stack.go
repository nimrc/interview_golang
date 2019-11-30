package common

/*
 * 栈 Stack LIFO
 */

type Stack struct {
	list DLinkedList
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(value int) {
	stack.list.Lpush(value)
}

func (stack *Stack) Pop() int {
	return stack.list.Lpop()
}

func (stack *Stack) Length() int {
	return stack.list.Lenth()
}

func (stack *Stack) IsEmpty() bool {
	return stack.list.IsEmpty()
}
