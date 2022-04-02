package linked_queue

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

//IsEmpty return true if the queue is empty.
func (l *LinkedQueue) IsEmpty() bool {
	return l.size == 0
}

//Len return the size of the queue
func (l *LinkedQueue) Len() int {
	return l.size
}
