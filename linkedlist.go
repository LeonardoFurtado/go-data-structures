package go_data_structures

//
//
// Linked List Data Structure
//
//┌─────┬┐    ┌─────┬┐    ┌─────┬┐
//│     │┼───►│     │┼───►│     ││
//└─────┴┘    └─────┴┘    └─────┴┘

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) AddFirst(val int) {
	newNode := Node{val, l.head}
	l.head = &newNode
	l.len++
}

// AddLast adds a value val to the end of the linked list
func (l *LinkedList) AddLast(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	node := l.head
	for i := 0; i < l.len; i++ {
		if node.next == nil {
			node.next = &n
			l.len++
			return
		}
		node = node.next
	}
}

// Insert adds a value val at position
func (l *LinkedList) Insert(position, val int) {
	if position > l.len+1 {
		panic("Position not allowed")
	}
	if position == 1 {
		l.AddFirst(val)
		return
	}
	n := Node{}
	n.value = val
	node := l.head
	for i := 1; i < position-1; i++ {
		node = node.next
	}
	n.next = node.next
	node.next = &n
	l.len++
}

// RemoveFirst removes a value at head
func (l *LinkedList) RemoveFirst() {
	if l.len == 0 {
		panic("The list is empty.")
	}
	l.head = l.head.next
	l.len--
}

// RemoveAt removes a value at position
func (l *LinkedList) RemoveAt(position int) {
	if position < 1 || position > l.len {
		panic("Index is not valid")
	}

	if position == 1 {
		l.head = l.head.next
		l.len--
	}
	node := l.head

	for i := 1; i < position-1; i++ {
		node = node.next
	}
	node.next = node.next.next
	l.len--
}

// IsEmpty checks if the list has any Node
func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

// IndexOf retrieve the position of the first value found
func (l *LinkedList) IndexOf(value int) (result int) {
	node := l.head
	for i := 1; i <= l.len; i++ {
		if node.value == value {
			result = i
			break
		}
		node = node.next
	}
	return result
}

// Size does what you think it does
func (l *LinkedList) Size() int {
	return l.len
}
