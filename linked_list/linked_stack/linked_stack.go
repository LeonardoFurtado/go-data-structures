package linked_stack

type Node struct {
	element int
	next    *Node
}

type LinkedStack struct {
	head *Node
	size int
}

func (l *LinkedStack) Push(value int) {
	newNode := Node{value, l.head}
	l.head = &newNode
	l.size++
}

func (l *LinkedStack) Pop() int {
	popValue := l.head.element
	l.head = l.head.next
	l.size--
	return popValue
}

func (l *LinkedStack) Top() int {
	return l.head.element
}

func (l *LinkedStack) Len() int {
	return l.size
}

func (l *LinkedStack) IsEmpty() bool {
	return l.size == 0
}
