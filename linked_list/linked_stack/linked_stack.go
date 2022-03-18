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
