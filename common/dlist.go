package common

/*
 * 双端链表 Double linked list
 */

type DListNode struct {
	Data int
	Prev *DListNode
	Next *DListNode
}

type DLinkedList struct {
	Head *DListNode
	Tail *DListNode
	Len  int
}

func NewDLinkedList() *DLinkedList {
	list := &DLinkedList{
		Head: nil,
		Tail: nil,
		Len:  0,
	}

	return list
}

func (list *DLinkedList) Lpush(value int) {
	node := &DListNode{Data: value, Prev: nil, Next: nil,}

	if list.Head == nil {
		list.Head = node
		list.Tail = list.Head
	} else {
		list.Head.Prev = node
		node.Next = list.Head
		list.Head = node
	}

	list.Len++
}

func (list *DLinkedList) Rpush(value int) {
	node := &DListNode{Data: value, Prev: nil, Next: nil,}

	if list.Tail == nil {
		list.Tail = node
		list.Head = list.Tail
	} else {
		list.Tail.Next = node
		node.Prev = list.Tail
		list.Tail = node
	}

	list.Len++
}

func (list *DLinkedList) Lpop() int {
	if list.Len == 0 {
		panic("list is empty")
	}

	node := list.Head
	list.Head = node.Next

	if node.Next != nil {
		node.Next.Prev = nil
		node.Next = nil
	} else {
		list.Tail = node.Next
	}

	list.Len--

	return node.Data
}

func (list *DLinkedList) Rpop() int {
	if list.Len == 0 {
		panic("list is empty")
	}

	node := list.Tail
	list.Tail = node.Prev

	if node.Prev != nil {
		node.Prev.Next = nil
		node.Prev = nil
	} else {
		list.Head = node.Prev
	}

	list.Len--

	return node.Data
}

func (list *DLinkedList) Lenth() int {
	return list.Len
}

func (list *DLinkedList) IsEmpty() bool {
	return list.Len == 0
}
