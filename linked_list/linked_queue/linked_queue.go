package linked_queue

import "errors"

/*
FIFO queue implementation using a singly linked list for storage.
*/

type Node struct {
	element int
	next    *Node
}

type LinkedQueue struct {
	head *Node
	tail *Node
	size int
}

//Enqueue add an element to the back of queue.
func (l *LinkedQueue) Enqueue(element int) {
	newest := &Node{element, nil}
	if l.IsEmpty() {
		l.head = newest
	} else {
		l.tail.next = newest
	}
	l.tail = newest
	l.size++
}

//Dequeue remove and return the ﬁrst element of the queue (i.e., FIFO).
func (l *LinkedQueue) Dequeue() (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	answer := l.head.element
	l.head = l.head.next
	l.size -= 1
	if l.IsEmpty() {
		l.tail = nil
	}
	return answer, nil
}

//IsEmpty return true if the queue is empty.
func (l *LinkedQueue) IsEmpty() bool {
	return l.size == 0
}

//Len return the size of the queue
func (l *LinkedQueue) Len() int {
	return l.size
}

//First return the element of the head node
func (l *LinkedQueue) First() (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return l.head.element, nil
}
