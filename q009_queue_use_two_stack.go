package main

import (
	"fmt"
	. "github.com/fyibmsd/interview/common"
)

/*
 * Q: 两个栈实现一个队列
 */
type CQueue struct {
	S1  Stack
	S2  Stack
	Len int
}

func NewCQueue() *CQueue {
	return &CQueue{}
}

func (queue *CQueue) Offer(value int) {
	queue.S1.Push(value)
	queue.Len++
}

func (queue *CQueue) Poll() int {
	if queue.Len == 0 {
		panic("queue is empty")
	}

	if queue.S2.IsEmpty() {
		for !queue.S1.IsEmpty() {
			queue.S2.Push(queue.S1.Pop())
		}
	}

	value := queue.S2.Pop()
	queue.Len--

	return value
}

func (queue *CQueue) IsEmpty() bool {
	return queue.Len == 0
}

func main() {
	values := []int{1, 3, 5, 7, 9}

	queue := NewCQueue()

	for _, v := range values {
		queue.Offer(v)
	}

	for !queue.IsEmpty() {
		fmt.Println(queue.Poll())
	}

}
