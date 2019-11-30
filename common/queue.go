package common

/*
 * 队列 Queue FIFO
 */

type Queue struct {
	list DLinkedList
}

func NewQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) Offer(value int) {
	queue.list.Lpush(value)
}

func (queue *Queue) Pop() int {
	return queue.list.Rpop()
}

func (queue *Queue) Length() int {
	return queue.list.Lenth()
}

func (queue *Queue) IsEmpty() bool {
	return queue.list.IsEmpty()
}
